package gamedata

import (
	"elichika/dictionary"
	"elichika/model"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type TrainingTreeCellItemSet struct {
	// from m_training_tree_cell_item_set
	ID        int             `xorm:"pk 'id'"`
	Resources []model.Content `xorm:"-"`
}

func (set *TrainingTreeCellItemSet) populate(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	err := masterdata_db.Table("m_training_tree_cell_item_set").Where("id = ?", set.ID).Find(&set.Resources)
	utils.CheckErr(err)
}

func loadTrainingTreeCellItemSet(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading TrainingCellItemSet")
	gamedata.TrainingTreeCellItemSet = make(map[int]*TrainingTreeCellItemSet)
	err := masterdata_db.Table("m_training_tree_cell_item_set").Find(gamedata.TrainingTreeCellItemSet)
	utils.CheckErr(err)
	for _, set := range gamedata.TrainingTreeCellItemSet {
		set.populate(gamedata, masterdata_db, serverdata_db, dictionary)
	}
}

func init() {
	addLoadFunc(loadTrainingTreeCellItemSet)
}