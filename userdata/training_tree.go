package userdata

import (
	"elichika/client"
	"elichika/generic"
	"elichika/utils"
)

// return the training tree for a card
func (session *Session) GetTrainingTree(cardMasterId int32) generic.List[client.UserCardTrainingTreeCell] {
	cells := generic.List[client.UserCardTrainingTreeCell]{}
	err := session.Db.Table("u_card_training_tree_cell").
		Where("user_id = ? AND card_master_id = ?", session.UserId, cardMasterId).Find(&cells.Slice)
	utils.CheckErr(err)
	return cells
}

// insert a training cell sets
func (session *Session) InsertTrainingTreeCells(cardMasterId int32, cells []client.UserCardTrainingTreeCell) {

	type Wrapper struct {
		CardMasterId int32 
		CellId      int32
		ActivatedAt int64
	}
	for _, cell := range cells {
		genericDatabaseInsert(session, "u_card_training_tree_cell", Wrapper{
			CardMasterId: cardMasterId, 
			CellId: cell.CellId,
			ActivatedAt: cell.ActivatedAt,
		})
	}
}