package account

import (
	"elichika/client"
	"elichika/gamedata"
	"elichika/model"
	"elichika/protocol/response"
	"elichika/userdata"
)

// we need to find a set of cards that is consistent with the card data
// in theory is equivelant to summing vectors in however many dimension necessary to describe a card training
// - IsAwakening
// - TrainingActivatedCellCount
// - MaxFreePassiveSkill
// - TrainingLife
// - TrainingAttack
// - TrainingDexterity
// - ActiveSkillLevel
// - PassiveSkillALevel
// - Side Story unlocked
// - Voice unlocked
// - Suit unlocked
// we can iterate over all possible tree, the complexity is then lower bounded by the number of way to pick for the forks
// which come out to at least 6 * 5 * 5 * 5 * 5 * 6 * 2 * 2 * 3 * 3 * 3 * 4 * 4 * 4 * 4 * 4 * 4 for a festival card
// for each of the way to pick fork, we have to iterate over the middle
// in theory it's possible to get it done in like 100s assuming worst case, but that's too slow
// we use the following optimization:
// - first select the dimensions that we know whether we have to take a specific tile or not (ban-pick dimensions):
//   - IsAwakening (pick 0/1 tile out of 1)
//   - Side story unlocked (0/1 tile out of 1)
//   - Voice unlocked (2 sets with 0/1 tile out of 1)
// - secondly select the dimensions that are relatively simple
//   - MaxFreePassiveSkill (pick 0/1/2 tile of out of up to 2), at most 2C1 = 2 choices
//   - ActiveSkillLevel (pick 0-4 out of 4 tiles, or 0-6 out of 6 tiles). at most 6C3 = 20 choices
//   - PassiveSkillALevel ActiveSkillLevel. At most 20 choices
// - this narrow us down to at most 800 basic set of tiles, some of them are also inconsistent and can be skipped
// - then the rest of the params are small enough to be solved with dynamic programming
//   - because the tiles of the same type mostly have 2 or at most 3 different value, it's not that bad if we store the state in a map or something
// finally, if we can't find a correct set of tiles then we reset the card to completely new just to make sure things are consistent

const (
	BFDimensionMaxFreePassiveSkill = 0
	BFDimensionActiveSkillLevel    = 1
	BFDimensionPassiveSkillALevel  = 2
	BFDimensionCount               = 3

	DPDimensionTrainingActivatedCellCount = 0
	DPDimensionTrainingLife               = 1
	DPDimensionTrainingAttack             = 2
	DPDimensionTrainingDexterity          = 3
	DPDimensionCount                      = 4
)

type SolverNode struct {
	Id       int
	DPWeight [DPDimensionCount]int
	Parent   *SolverNode
	Children []*SolverNode

	IsPicked bool
	IsBanned bool
}

type TrainingTreeSolver struct {
	// can be reused if single threaded, for multi threading, each will have to have its own data
	HasNaviVoice        map[int]bool
	HasStorySide        map[int]bool
	HasSuit             map[int32]bool
	Node                [100]SolverNode // need at most 91
	NodeCount           int
	TrainingTree        *gamedata.TrainingTree
	TrainingTreeMapping *gamedata.TrainingTreeMapping
	TrainingTreeDesign  *gamedata.TrainingTreeDesign
	MasterCard          *gamedata.Card
	Cells               []model.TrainingTreeCell
	TimeStamp           int64
	Session             *userdata.Session
	Card                *client.UserCard

	BFNodes   [BFDimensionCount]([]*SolverNode)
	BFTarget  [BFDimensionCount]int
	BFCurrent [BFDimensionCount]int

	OperationStack [100]struct {
		Ptr      *bool
		OldValue bool
	}
	OperationCount int
}

func (solver *TrainingTreeSolver) PerformOperation(ptr *bool, newValue bool) {
	solver.OperationCount++
	solver.OperationStack[solver.OperationCount].Ptr = ptr
	solver.OperationStack[solver.OperationCount].OldValue = *ptr
	*ptr = newValue
}

func (solver *TrainingTreeSolver) UndoOperations(operationLeft int) {
	for ; solver.OperationCount > operationLeft; solver.OperationCount-- {
		*solver.OperationStack[solver.OperationCount].Ptr = solver.OperationStack[solver.OperationCount].OldValue
	}
}

func (solver *TrainingTreeSolver) MarkPicked(node *SolverNode) bool {
	// if a node is picked then its ancestors are also picked
	for ; node.Id != 0; node = node.Parent {
		if node.IsPicked {
			return true
		}
		if node.IsBanned {
			return false
		}
		solver.PerformOperation(&node.IsPicked, true)
	}
	return true
}
func (solver *TrainingTreeSolver) MarkBanned(node *SolverNode) bool {
	// we only need to make this one node because the checking will be done by the picked side
	if node.IsPicked {
		return false
	}
	if !node.IsBanned {
		solver.PerformOperation(&node.IsBanned, true)
	}
	return true
}
func (node *SolverNode) Prepare(solver *TrainingTreeSolver) {
	node.Parent = &solver.Node[solver.TrainingTreeDesign.Parent[node.Id]]
	node.Children = nil
	for _, child := range solver.TrainingTreeDesign.Children[node.Id] {
		node.Children = append(node.Children, &solver.Node[child])
	}
	node.IsPicked = false
	node.IsBanned = false
	for i := range node.DPWeight {
		node.DPWeight[i] = 0
	}
}
func (node *SolverNode) Populate(solver *TrainingTreeSolver) bool {
	// get the content
	cell := solver.TrainingTreeMapping.TrainingTreeCellContents[node.Id]
	if cell.RequiredGrade > int(solver.Card.Grade) {
		return solver.MarkBanned(node)
	}
	node.DPWeight[DPDimensionTrainingActivatedCellCount] = 1
	switch cell.TrainingTreeCellType {
	case 2: // params
		paramCell := solver.TrainingTree.TrainingTreeCardParams[cell.TrainingContentNo]
		switch paramCell.TrainingContentType {
		case 2: // stamina
			node.DPWeight[DPDimensionTrainingLife] += paramCell.Value
		case 3: // appeal
			node.DPWeight[DPDimensionTrainingAttack] += paramCell.Value
		case 4: // technique
			node.DPWeight[DPDimensionTrainingDexterity] += paramCell.Value
		default:
			panic("Unexpected training content type")
		}
		return true
	case 3: // voice
		naviActionId := solver.TrainingTree.NaviActionIds[cell.TrainingContentNo]
		if solver.HasNaviVoice[naviActionId] {
			return solver.MarkPicked(node)
		} else {
			return solver.MarkBanned(node)
		}
	case 4: // story cell
		storySideId, exist := solver.TrainingTree.TrainingTreeCardStorySides[11]
		if !exist {
			panic("story doesn't exist")
		}
		if solver.HasStorySide[storySideId] {
			return solver.MarkPicked(node)
		} else {
			return solver.MarkBanned(node)
		}
	case 5:
		// idolize
		if solver.Card.IsAwakening {
			return solver.MarkPicked(node)
		} else {
			return solver.MarkBanned(node)
		}
	case 6: // costume
		suitId := solver.TrainingTree.SuitMIds[cell.TrainingContentNo]
		if solver.HasSuit[suitId] {
			return solver.MarkPicked(node)
		} else {
			return solver.MarkBanned(node)
		}
	case 7: // skill
		solver.BFNodes[BFDimensionActiveSkillLevel] = append(solver.BFNodes[BFDimensionActiveSkillLevel], node)
	case 8: // insight
		solver.BFNodes[BFDimensionMaxFreePassiveSkill] = append(solver.BFNodes[BFDimensionMaxFreePassiveSkill], node)
	case 9: // ability
		solver.BFNodes[BFDimensionPassiveSkillALevel] = append(solver.BFNodes[BFDimensionPassiveSkillALevel], node)
	default:
		panic("Unknown cell type")
		return false
	}
	return true
}

func (solver *TrainingTreeSolver) LoadUserLogin(login *response.Login) {
	solver.HasNaviVoice = make(map[int]bool)
	solver.HasStorySide = make(map[int]bool)
	solver.HasSuit = make(map[int32]bool)
	for _, voice := range login.UserModel.UserVoiceByVoiceId.Objects {
		solver.HasNaviVoice[voice.NaviVoiceMasterId] = true
	}
	for _, storySide := range login.UserModel.UserStorySideById.Objects {
		solver.HasStorySide[storySide.StorySideMasterId] = true
	}
	for _, suit := range login.UserModel.UserSuitBySuitId.Objects {
		solver.HasSuit[suit.SuitMasterId] = true
	}
}

func (solver *TrainingTreeSolver) AddCell(cellId int) {
	if cellId == 0 {
		return
	}
	solver.Cells = append(solver.Cells,
		model.TrainingTreeCell{
			CardMasterId: int(solver.Card.CardMasterId),
			CellId:       cellId,
			ActivatedAt:  solver.TimeStamp})
}
