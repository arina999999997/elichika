package gamedata

import (
	"elichika/client"

	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type TrainingTreeCellItemSet struct {
	// from m_training_tree_cell_item_set
	Id        int32            `xorm:"pk 'id'"`
	Resources []client.Content `xorm:"-"`
}

func (set *TrainingTreeCellItemSet) populate(gamedata *Gamedata) {
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_training_tree_cell_item_set").Where("id = ?", set.Id).Find(&set.Resources)
	})
	utils.CheckErr(err)
}

func loadTrainingTreeCellItemSet(gamedata *Gamedata) {
	fmt.Println("Loading TrainingCellItemSet")
	gamedata.TrainingTreeCellItemSet = make(map[int32]*TrainingTreeCellItemSet)
	var err error
	gamedata.MasterdataDb.Do(func(session *xorm.Session) {
		err = session.Table("m_training_tree_cell_item_set").Find(gamedata.TrainingTreeCellItemSet)
	})
	utils.CheckErr(err)
	for _, set := range gamedata.TrainingTreeCellItemSet {
		set.populate(gamedata)
	}
}

func init() {
	addLoadFunc(loadTrainingTreeCellItemSet)
}
