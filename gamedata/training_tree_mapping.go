package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type TrainingTreeCellContent struct {
	// from m_training_tree_cell_content
	// ID int `xorm:"pk 'id'"`
	// CellID int `xorm:"pk 'cell_id'"`
	TrainingTreeCellType int `xorm:"'training_tree_cell_type'"`
	TrainingContentNo    int `xorm:"'training_content_no'"`
	// RequiredGrade int `xorm:"'required_grade'"`
	TrainingTreeCellItemSetMID *int                     `xorm:"'training_tree_cell_item_set_m_id'"`
	TrainingTreeCellItemSet    *TrainingTreeCellItemSet `xorm:"-"`

	SnsCoin int `xorm:"'sns_coin'"`
}

func (obj *TrainingTreeCellContent) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	obj.TrainingTreeCellItemSet = gamedata.TrainingTreeCellItemSet[*obj.TrainingTreeCellItemSetMID]
	obj.TrainingTreeCellItemSetMID = &gamedata.TrainingTreeCellItemSet[*obj.TrainingTreeCellItemSetMID].ID
}

type TrainingTreeMapping struct {
	// from m_training_tree_mapping
	ID                         int                       `xorm:"pk 'id'"`
	TrainingTreeCellContentMID int                       `xorm:"'training_tree_cell_content_m_id'"`
	TrainingTreeCellContents   []TrainingTreeCellContent `xorm:"-"` // 0 indexed
}

func (tree_mapping *TrainingTreeMapping) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	err := masterdata_db.Table("m_training_tree_cell_content").Where("id = ?", tree_mapping.TrainingTreeCellContentMID).
		OrderBy("cell_id").Find(&tree_mapping.TrainingTreeCellContents)
	utils.CheckErr(err)
	for i := range tree_mapping.TrainingTreeCellContents {
		tree_mapping.TrainingTreeCellContents[i].populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func loadTrainingTreeMapping(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading TrainingMapping")
	gamedata.TrainingTreeMapping = make(map[int]*TrainingTreeMapping)
	err := masterdata_db.Table("m_training_tree_mapping").Find(&gamedata.TrainingTreeMapping)
	utils.CheckErr(err)
	for _, tree_mapping := range gamedata.TrainingTreeMapping {
		tree_mapping.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadTrainingTreeMapping)
	addPrequisite(loadTrainingTreeMapping, loadTrainingTreeCellItemSet)
}
