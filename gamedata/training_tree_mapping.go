package gamedata

import (
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type TrainingTreeCellContent struct {
	// from m_training_tree_cell_content
	// Id int `xorm:"'id'"`
	CellId                     int32                    `xorm:"'cell_id'"`
	TrainingTreeCellType       int32                    `xorm:"'training_tree_cell_type'" enum:"TrainingTreeCellType"`
	TrainingContentNo          int32                    `xorm:"'training_content_no'"`
	RequiredGrade              int32                    `xorm:"'required_grade'"`
	TrainingTreeCellItemSetMId *int32                   `xorm:"'training_tree_cell_item_set_m_id'"`
	TrainingTreeCellItemSet    *TrainingTreeCellItemSet `xorm:"-"`

	SnsCoin int `xorm:"'sns_coin'"`
}

func (obj *TrainingTreeCellContent) populate(gamedata *Gamedata) {
	obj.TrainingTreeCellItemSet = gamedata.TrainingTreeCellItemSet[*obj.TrainingTreeCellItemSetMId]
	obj.TrainingTreeCellItemSetMId = &gamedata.TrainingTreeCellItemSet[*obj.TrainingTreeCellItemSetMId].Id
}

type TrainingTreeMapping struct {
	// from m_training_tree_mapping
	Id                         int32                     `xorm:"pk 'id'"`
	TrainingTreeCellContentMId int32                     `xorm:"'training_tree_cell_content_m_id'"`
	TrainingTreeCellContents   []TrainingTreeCellContent `xorm:"-"` // 0 indexed
	TrainingTreeDesignMId      *int32                    `xorm:"'training_tree_design_m_id'"`
	TrainingTreeDesign         *TrainingTreeDesign       `xorm:"-"`
}

func (treeMapping *TrainingTreeMapping) populate(gamedata *Gamedata) {
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_training_tree_cell_content").Where("id = ?", treeMapping.TrainingTreeCellContentMId).OrderBy("cell_id").Find(&treeMapping.TrainingTreeCellContents)
	})
	utils.CheckErr(err)
	for i := range treeMapping.TrainingTreeCellContents {
		treeMapping.TrainingTreeCellContents[i].populate(gamedata)
	}
	treeMapping.TrainingTreeDesign = gamedata.TrainingTreeDesign[*treeMapping.TrainingTreeDesignMId]
	treeMapping.TrainingTreeDesignMId = &treeMapping.TrainingTreeDesign.Id
}

func loadTrainingTreeMapping(gamedata *Gamedata) {
	fmt.Println("Loading TrainingMapping")
	gamedata.TrainingTreeMapping = make(map[int32]*TrainingTreeMapping)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_training_tree_mapping").Find(&gamedata.TrainingTreeMapping)
	})
	utils.CheckErr(err)
	for _, treeMapping := range gamedata.TrainingTreeMapping {
		treeMapping.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadTrainingTreeMapping)
	addPrequisite(loadTrainingTreeMapping, loadTrainingTreeCellItemSet)
	addPrequisite(loadTrainingTreeMapping, loadTrainingTreeDesign)
}
