package user_training_tree

import (
	"elichika/client"
	"elichika/userdata"
)

// insert a training cell sets
func InsertUserTrainingTreeCells(session *userdata.Session, cardMasterId int32, cells []client.UserCardTrainingTreeCell) {
	type Wrapper struct {
		CardMasterId int32
		CellId       int32
		ActivatedAt  int64
	}
	for _, cell := range cells {
		userdata.GenericDatabaseInsert(session, "u_card_training_tree_cell", Wrapper{
			CardMasterId: cardMasterId,
			CellId:       cell.CellId,
			ActivatedAt:  cell.ActivatedAt,
		})
	}
}
