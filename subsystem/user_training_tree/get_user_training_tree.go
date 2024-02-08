package user_training_tree

import (
	"elichika/client"
	"elichika/generic"
	"elichika/utils"
	"elichika/userdata"
)

// return the training tree for a card
func GetUserTrainingTree(session *userdata.Session, cardMasterId int32) generic.List[client.UserCardTrainingTreeCell] {
	cells := generic.List[client.UserCardTrainingTreeCell]{}
	err := session.Db.Table("u_card_training_tree_cell").
		Where("user_id = ? AND card_master_id = ?", session.UserId, cardMasterId).Find(&cells.Slice)
	utils.CheckErr(err)
	return cells
}