package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

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
type TrainingTreeDesign struct {
	// from m_training_tree_design
	Id        int
	CellCount int
	Parent    []int
	Children  []([]int)
	// ParentBranchType []int
}

func loadTrainingTreeDesign(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading TrainingTreeDesign")
	type TrainingTreeDesignCell struct {
		DesignId         int `xorm:"'id'"`
		CellId           int `xorm:"'cell_id'"`
		ParentCellId     int `xorm:"'parent_cell_id'"`
		ParentBranchType int `xorm:"'parent_branch_type'"`
	}
	cells := []TrainingTreeDesignCell{}
	err := masterdata_db.Table("m_training_tree_design").Find(&cells)
	utils.CheckErr(err)
	gamedata.TrainingTreeDesign = make(map[int]*TrainingTreeDesign)
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
		for i := 0; i < design.CellCount; i++ {
			design.Parent = append(design.Parent, 0)
			design.Children = append(design.Children, []int{})
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