package userdata

import (
	"elichika/model"
	"elichika/utils"

	"fmt"
)

// return the training tree for a card
func (session *Session) GetTrainingTree(cardMasterID int) []model.TrainingTreeCell {
	cells := []model.TrainingTreeCell{}
	err := session.Db.Table("u_training_tree_cell").
		Where("user_id = ? AND card_master_id = ?", session.UserStatus.UserID, cardMasterID).Find(&cells)
	utils.CheckErr(err)
	return cells
}

// insert a training cell sets
func (session *Session) InsertTrainingTreeCells(cells []model.TrainingTreeCell) {
	session.UserTrainingTreeCellDiffs = append(session.UserTrainingTreeCellDiffs, cells...)
}

func finalizeTrainingTree(session *Session) {
	n := len(session.UserTrainingTreeCellDiffs)
	// this is not for speed (alone), it's to avoid SQL freaking out because there's too many params
	step := 5000
	for begin := 0; begin < n; begin += step {
		end := begin + step
		if end > n {
			end = n
		}
		affected, err := session.Db.Table("u_training_tree_cell").AllCols().Insert(session.UserTrainingTreeCellDiffs[begin:end])
		utils.CheckErr(err)
		fmt.Println("Inserted ", affected, " training cells")
	}
}

func init() {
	addFinalizer(finalizeTrainingTree)
}
