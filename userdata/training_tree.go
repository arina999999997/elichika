package userdata

import (
	"elichika/model"
	"elichika/utils"
)

// return the training tree for a card
func (session *Session) GetTrainingTree(cardMasterId int) []model.TrainingTreeCell {
	cells := []model.TrainingTreeCell{}
	err := session.Db.Table("u_training_tree_cell").
		Where("user_id = ? AND card_master_id = ?", session.UserId, cardMasterId).Find(&cells)
	utils.CheckErr(err)
	return cells
}

// insert a training cell sets
func (session *Session) InsertTrainingTreeCells(cells []model.TrainingTreeCell) {
	session.UserTrainingTreeCellDiffs = append(session.UserTrainingTreeCellDiffs, cells...)
}

func finalizeTrainingTree(session *Session) {
	for _, cell := range session.UserTrainingTreeCellDiffs {
		genericDatabaseInsert(session, "u_training_tree_cell", cell)
	}
}

func init() {
	addFinalizer(finalizeTrainingTree)
}
