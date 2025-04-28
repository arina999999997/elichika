package user_training_tree

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"
)

// return the training tree for a card
func GetUserTrainingTree(session *userdata.Session, cardMasterId int32) generic.List[client.UserCardTrainingTreeCell] {
	storedCells := []client.UserCardTrainingTreeCell{}
	err := session.Db.Table("u_card_training_tree_cell").
		Where("user_id = ? AND card_master_id = ?", session.UserId, cardMasterId).Find(&storedCells)
	utils.CheckErr(err)
	cells := generic.List[client.UserCardTrainingTreeCell]{
		Slice: session.Gamedata.TrainingTree[cardMasterId].Design().Uncompress(storedCells),
	}
	return cells
}
