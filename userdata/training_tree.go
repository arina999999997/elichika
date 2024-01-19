package userdata

import (
	"elichika/client"
	"elichika/generic"
	"elichika/model"
	"elichika/utils"
)

// return the training tree for a card
func (session *Session) GetTrainingTree(cardMasterId int32) generic.List[client.UserCardTrainingTreeCell] {
	cells := generic.List[client.UserCardTrainingTreeCell]{}
	err := session.Db.Table("u_training_tree_cell").
		Where("user_id = ? AND card_master_id = ?", session.UserId, cardMasterId).Find(&cells.Slice)
	utils.CheckErr(err)
	return cells
}

// insert a training cell sets
func (session *Session) InsertTrainingTreeCells(cardMasterId int32, cells []client.UserCardTrainingTreeCell) {
	for _, cell := range cells {
		session.UserTrainingTreeCellDiffs = append(session.UserTrainingTreeCellDiffs, model.TrainingTreeCell{
			CardMasterId: int(cardMasterId),
			CellId:       int(cell.CellId),
			ActivatedAt:  cell.ActivatedAt,
		})
	}
}

func finalizeTrainingTree(session *Session) {
	for _, cell := range session.UserTrainingTreeCellDiffs {
		genericDatabaseInsert(session, "u_training_tree_cell", cell)
	}
}

func init() {
	addFinalizer(finalizeTrainingTree)
}
