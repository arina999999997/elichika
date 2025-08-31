package gamedata

import (
	"elichika/client"
	"elichika/log"
	"elichika/utils"

	"xorm.io/xorm"
)

// The training tree is just a tree (in the computer science sense)
// Which mean it is fully represented by the parent array of each node
// We also produce the children array of each node for convinience when used
// ParentBranchType just dictate how the tree should be drawn.
// - 3 is the root node
// - 100 is leveled (horizontally the same)
// - 101 is going up
// - 102 is going down
// Children will have the main child (ParentBranchType = 100) as the first item, otherwise it doesn't matter what come first
// Assume the following:
// - Parent node cell id is always smaller than node cell id
// - Node 0 is the special root and is often ignored
type TrainingTreeDesign struct {
	// from m_training_tree_design
	Id        int32
	CellCount int32
	Parent    []int32
	Children  []([]int32)
	// ParentBranchType []int32
}

func (tree *TrainingTree) Design() *TrainingTreeDesign {
	return tree.TrainingTreeMapping.TrainingTreeDesign
}

// compressed form of the training tree:
// - if a node is unlocked along side its child(ren), then only store the child(ren)
//   - and we do this recursively, so only the final unlocked nodes are stored
// - then we have some nodes that are stored in the training tree along side their unlocking times:
//   - if the unlocking time for a node is not present, then it must have been unlocked along side its children, recursively
// - in the usual usage one might see in a private server, unlocking everything at once would only store around 1/5 of the data
//
// TODO(extra): the following schemes might offer better compression over all:
// - map from unlocked times to bitsets of cell
// - negative storage: store the nodes that are ready to be unlocked but not unlocked, and store the unlocked time of the node closest to the root:
//   - this has the avantage that a tree unlocked all at once only consume one cell
//   - A tree unlocked once each time a limit break is gained only consume 6 cells
// - the above would requires some extra storage that is not compatible with the naive storage method, so some changes will be needed

// Compress a slice of cells using whatever method we have
// assume the slice contain correct cells in the naive form
func (design *TrainingTreeDesign) Compress(cells []client.UserCardTrainingTreeCell) []client.UserCardTrainingTreeCell {
	activateTimes := make([]int64, int(design.CellCount))
	for _, cell := range cells {
		activateTimes[int(cell.CellId)] = cell.ActivatedAt
	}
	for _, cell := range cells {
		cellId := int(design.Parent[int(cell.CellId)])
		activatedAt := cell.ActivatedAt
		for cellId > 0 {
			// mark all parents that share the activate time as unnecessary
			if activateTimes[cellId] != activatedAt {
				break
			}
			activateTimes[cellId] = 0
			cellId = int(design.Parent[int(cellId)])
		}
	}
	returnCells := []client.UserCardTrainingTreeCell{}
	for _, cell := range cells {
		if activateTimes[int(cell.CellId)] == 0 {
			continue
		}
		returnCells = append(returnCells, cell)
	}
	return returnCells
}

// Uncompress a slice of cells using whatever method we have
// the slice can contain extra cells
func (design *TrainingTreeDesign) Uncompress(cells []client.UserCardTrainingTreeCell) []client.UserCardTrainingTreeCell {
	activateTimes := make([]int64, int(design.CellCount))
	for i := 0; i < int(design.CellCount); i++ {
		activateTimes[i] = (1 << 63) - 1
	}
	for _, cell := range cells {
		activateTimes[int(cell.CellId)] = cell.ActivatedAt
	}
	// some tree dp to store the data
	for i := design.CellCount - 1; i > 1; i-- {
		if activateTimes[design.Parent[i]] > activateTimes[i] {
			activateTimes[design.Parent[i]] = activateTimes[i]
		}
	}
	returnCells := []client.UserCardTrainingTreeCell{}
	for i := 1; i < int(design.CellCount); i++ {
		if activateTimes[i] == (1<<63)-1 {
			continue
		}
		returnCells = append(returnCells, client.UserCardTrainingTreeCell{
			CellId:      int32(i),
			ActivatedAt: activateTimes[i],
		})
	}
	return returnCells
}

func loadTrainingTreeDesign(gamedata *Gamedata) {
	log.Println("Loading TrainingTreeDesign")
	type TrainingTreeDesignCell struct {
		DesignId         int32 `xorm:"'id'"`
		CellId           int32 `xorm:"'cell_id'"`
		ParentCellId     int32 `xorm:"'parent_cell_id'"`
		ParentBranchType int32 `xorm:"'parent_branch_type'"`
	}
	cells := []TrainingTreeDesignCell{}
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_training_tree_design").Find(&cells)
	})
	utils.CheckErr(err)
	gamedata.TrainingTreeDesign = make(map[int32]*TrainingTreeDesign)
	for _, cell := range cells {
		_, exist := gamedata.TrainingTreeDesign[cell.DesignId]
		if !exist {
			gamedata.TrainingTreeDesign[cell.DesignId] = &TrainingTreeDesign{
				Id: cell.DesignId,
			}
		}
		gamedata.TrainingTreeDesign[cell.DesignId].CellCount++
	}
	for _, design := range gamedata.TrainingTreeDesign {
		for i := int32(0); i < design.CellCount; i++ {
			design.Parent = append(design.Parent, 0)
			design.Children = append(design.Children, []int32{})
		}
	}
	for _, cell := range cells {
		design := gamedata.TrainingTreeDesign[cell.DesignId]
		design.Parent[cell.CellId] = cell.ParentCellId
		if cell.ParentBranchType == 100 {
			design.Children[cell.ParentCellId] = append(design.Children[cell.ParentCellId], cell.CellId)
		}
	}
	for _, cell := range cells {
		design := gamedata.TrainingTreeDesign[cell.DesignId]
		if cell.ParentBranchType != 100 {
			design.Children[cell.ParentCellId] = append(design.Children[cell.ParentCellId], cell.CellId)
		}
	}
}

func init() {
	addLoadFunc(loadTrainingTreeDesign)
}
