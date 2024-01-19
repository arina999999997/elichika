package account

import (
	"elichika/client"
	"elichika/userdata"

	"fmt"
)

func (solver *TrainingTreeSolver) SolveCard(session *userdata.Session, card client.UserCard) {
	solver.Cells = []client.UserCardTrainingTreeCell{}
	solver.Session = session
	solver.Card = &card
	solver.MasterCard = session.Gamedata.Card[card.CardMasterId]
	solver.TrainingTree = solver.MasterCard.TrainingTree
	solver.TrainingTreeMapping = solver.TrainingTree.TrainingTreeMapping
	solver.TrainingTreeDesign = solver.TrainingTreeMapping.TrainingTreeDesign
	solver.TimeStamp = int64(card.AcquiredAt)
	solver.NodeCount = solver.TrainingTreeDesign.CellCount - 1 // not counting the starting node id 0
	if card.IsAllTrainingActivated {                           // if maxed out then we don't need to solve
		for _, cell := range solver.TrainingTreeMapping.TrainingTreeCellContents {
			solver.AddCell(cell.CellId)
		}
	} else if card.TrainingActivatedCellCount == 0 {
		// entirely new card, no need to do anything
	} else if !solver.SolveForTileSet() { // otherwise we solve for a possible set of tiles
		fmt.Println("Solving failed for card", card.CardMasterId, ", reseting to default")
		session.UserModel.UserCardByCardId.Set(card.CardMasterId, client.UserCard{
			CardMasterId:        card.CardMasterId,
			Level:               card.Level,
			MaxFreePassiveSkill: solver.MasterCard.PassiveSkillSlot,
			Grade:               card.Grade, // check this for new card
			ActiveSkillLevel:    1,
			PassiveSkillALevel:  1,
			PassiveSkillBLevel:  1,
			PassiveSkillCLevel:  1,
			AcquiredAt:          card.AcquiredAt,
			IsNew:               true,
		})
	} // else {
	// fmt.Println("Found solution for card", card.CardMasterId)
	// }
	session.InsertTrainingTreeCells(card.CardMasterId, solver.Cells)
	if int32(len(solver.Cells)) != card.TrainingActivatedCellCount {
		panic(fmt.Sprint("wrong amount of cell, card master id: ", card.CardMasterId))
	}
	// update stat for this member
	userMember := session.GetMember(solver.MasterCard.Member.Id)
	userMember.OwnedCardCount++
	if card.IsAllTrainingActivated {
		userMember.AllTrainingCardCount++
	}
	session.UpdateMember(userMember)
}

func (solver *TrainingTreeSolver) SolveForTileSet() bool {
	// setup
	solver.OperationCount = 0
	for i := range solver.BFNodes {
		solver.BFNodes[i] = []*SolverNode{}
	}
	for i := 1; i <= solver.NodeCount; i++ {
		solver.Node[i].Id = i
	}
	solver.MarkPicked(&solver.Node[0])

	for i := 1; i <= solver.NodeCount; i++ {
		solver.Node[i].Prepare(solver)
	}
	for i := 1; i <= solver.NodeCount; i++ {
		if !solver.Node[i].Populate(solver) {
			// failed to pick a consistent set at the ban-pick phase
			return false
		}
	}
	// brute force smaller stuff phase, prepare the BF target
	solver.BFTarget[BFDimensionActiveSkillLevel] = int(solver.Card.ActiveSkillLevel - 1)
	solver.BFTarget[BFDimensionMaxFreePassiveSkill] = int(solver.Card.MaxFreePassiveSkill - solver.MasterCard.PassiveSkillSlot)
	solver.BFTarget[BFDimensionPassiveSkillALevel] = int(solver.Card.PassiveSkillALevel - 1)
	for i := range solver.BFCurrent {
		solver.BFCurrent[i] = 0
	}
	return solver.BruteForce(0, 0)
}

func (solver *TrainingTreeSolver) BruteForce(dim, item int) bool {
	if dim == BFDimensionCount {
		// brute force successful, we need to solve for stats with this set
		return solver.DynamicProgramming()
	}
	if item == len(solver.BFNodes[dim]) { // already done picking this dimension, pick for the next one
		return solver.BruteForce(dim+1, 0)
	}
	// if we can pick this item, then we have to mark it as picked
	if solver.BFCurrent[dim]+1 <= solver.BFTarget[dim] {
		backup := solver.OperationCount
		if solver.MarkPicked(solver.BFNodes[dim][item]) {
			solver.BFCurrent[dim]++
			if solver.BruteForce(dim, item+1) {
				return true
			}
			solver.BFCurrent[dim]--
		}
		solver.UndoOperations(backup)
	}
	// if we can skip this item, then we have to mark it as banned
	if solver.BFTarget[dim] <= solver.BFCurrent[dim]+len(solver.BFNodes[dim])-item-1 {
		backup := solver.OperationCount
		if solver.MarkBanned(solver.BFNodes[dim][item]) {
			if solver.BruteForce(dim, item+1) {
				return true
			}
		}
		solver.UndoOperations(backup)
	}
	return false
}

func (solver *TrainingTreeSolver) DynamicProgramming() bool {
	// iterate over the central path, then we can produce side path chains
	// we can take a prefix of the chains for dynamic programming

	state := [DPDimensionCount]int{int(solver.Card.TrainingActivatedCellCount),
		int(solver.Card.TrainingLife), int(solver.Card.TrainingAttack), int(solver.Card.TrainingDexterity)}
	wantedState := [DPDimensionCount]int{}
	for i := 1; i <= solver.NodeCount; i++ {
		if solver.Node[i].IsPicked {
			for j := range state {
				state[j] -= solver.Node[i].DPWeight[j]
			}
		}
	}

	exploredStates := map[[DPDimensionCount]int]([]*SolverNode){}
	exploredStates[state] = []*SolverNode{}
	centerNode := &solver.Node[1]
	solution, exists := exploredStates[wantedState]
	if exists {
		goto solutionFound
	}

	for !centerNode.IsBanned {
		if !centerNode.IsPicked {
			// add this node to every existing dp solution
			newExplored := map[[DPDimensionCount]int]([]*SolverNode){}
			for state, required := range exploredStates {
				nState := state
				for i := range nState {
					nState[i] -= centerNode.DPWeight[i]
					if nState[i] < 0 {
						goto nextStateCenter
					}
				}
				newExplored[nState] = required
			nextStateCenter:
			}
			solution, exists = newExplored[wantedState]
			if exists {
				goto solutionFound
			}
			exploredStates = newExplored
			if len(exploredStates) == 0 {
				break
			}
		}
		for i := range centerNode.Children {
			if i > 0 {
				sideNode := centerNode.Children[i]
				// produce a chain, skip things while sideNode is already picked
				for (sideNode != nil) && (sideNode.IsPicked) {
					if len(sideNode.Children) > 0 {
						sideNode = sideNode.Children[0]
					} else {
						sideNode = nil
					}
				}
				if (sideNode == nil) || sideNode.IsBanned { // there's no chain
					continue
				}
				chain := []*SolverNode{}
				for (sideNode != nil) && (!sideNode.IsBanned) {
					chain = append(chain, sideNode)
					if len(sideNode.Children) > 0 {
						sideNode = sideNode.Children[0]
					} else {
						sideNode = nil
					}
				}
				// update existing dp with this chain
				newExplored := map[[DPDimensionCount]int]([]*SolverNode){}
				for state, requiredNodes := range exploredStates {
					nState := state
					for _, lastNode := range chain {
						for i := range nState {
							nState[i] -= lastNode.DPWeight[i]
							if nState[i] < 0 {
								goto nextStateChain
							}
						}
						var nRequiredNodes []*SolverNode
						nRequiredNodes = append(nRequiredNodes, requiredNodes...)
						nRequiredNodes = append(nRequiredNodes, lastNode)
						newExplored[nState] = nRequiredNodes
					}
				nextStateChain:
				}
				solution, exists = newExplored[wantedState]
				if exists {
					goto solutionFound
				}
				for k, v := range newExplored {
					exploredStates[k] = v
				}
			}
		}
		if len(centerNode.Children) == 0 {
			break
		}
		centerNode = centerNode.Children[0]
	}
solutionFound:
	;
	if exists {
		solution = append(solution, centerNode)
		for _, node := range solution {
			if !solver.MarkPicked(node) {
				panic("wrong logic")
			}
		}
		for i := 1; i <= solver.NodeCount; i++ {
			if solver.Node[i].IsPicked {
				solver.AddCell(solver.Node[i].Id)
			}
		}
	}
	return exists
}
