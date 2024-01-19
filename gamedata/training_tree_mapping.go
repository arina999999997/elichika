package gamedata

import (
	"elichika/dictionary"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type TrainingTreeCellContent struct {
	// from m_training_tree_cell_content
	// Id int `xorm:"'id'"`
	CellId                     int                      `xorm:"'cell_id'"`
	TrainingTreeCellType       int                      `xorm:"'training_tree_cell_type'" enum:"TrainingTreeCellType"`
	TrainingContentNo          int                      `xorm:"'training_content_no'"`
	RequiredGrade              int                      `xorm:"'required_grade'"`
	TrainingTreeCellItemSetMId *int                     `xorm:"'training_tree_cell_item_set_m_id'"`
	TrainingTreeCellItemSet    *TrainingTreeCellItemSet `xorm:"-"`

	SnsCoin int `xorm:"'sns_coin'"`
}

func (obj *TrainingTreeCellContent) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	obj.TrainingTreeCellItemSet = gamedata.TrainingTreeCellItemSet[*obj.TrainingTreeCellItemSetMId]
	obj.TrainingTreeCellItemSetMId = &gamedata.TrainingTreeCellItemSet[*obj.TrainingTreeCellItemSetMId].Id
}

type TrainingTreeMapping struct {
	// from m_training_tree_mapping
	Id                         int                       `xorm:"pk 'id'"`
	TrainingTreeCellContentMId int                       `xorm:"'training_tree_cell_content_m_id'"`
	TrainingTreeCellContents   []TrainingTreeCellContent `xorm:"-"` // 0 indexed
	TrainingTreeDesignMId      *int                      `xorm:"'training_tree_design_m_id'"`
	TrainingTreeDesign         *TrainingTreeDesign       `xorm:"-"`
}

func (treeMapping *TrainingTreeMapping) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	err := masterdata_db.Table("m_training_tree_cell_content").Where("id = ?", treeMapping.TrainingTreeCellContentMId).
		OrderBy("cell_id").Find(&treeMapping.TrainingTreeCellContents)
	utils.CheckErr(err)
	for i := range treeMapping.TrainingTreeCellContents {
		treeMapping.TrainingTreeCellContents[i].populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
	treeMapping.TrainingTreeDesign = gamedata.TrainingTreeDesign[*treeMapping.TrainingTreeDesignMId]
	treeMapping.TrainingTreeDesignMId = &treeMapping.TrainingTreeDesign.Id
}

func loadTrainingTreeMapping(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading TrainingMapping")
	gamedata.TrainingTreeMapping = make(map[int]*TrainingTreeMapping)
	err := masterdata_db.Table("m_training_tree_mapping").Find(&gamedata.TrainingTreeMapping)
	utils.CheckErr(err)
	for _, treeMapping := range gamedata.TrainingTreeMapping {
		treeMapping.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadTrainingTreeMapping)
	addPrequisite(loadTrainingTreeMapping, loadTrainingTreeCellItemSet)
	addPrequisite(loadTrainingTreeMapping, loadTrainingTreeDesign)
}
