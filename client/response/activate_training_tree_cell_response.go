package response

import (
	"elichika/client"
	"elichika/generic"
)

type ActivateTrainingTreeCellResponse struct {
	UserCardTrainingTreeCellList generic.List[client.UserCardTrainingTreeCell] `json:"user_card_training_tree_cell_list"` // is actually named _UserCardTrainingTreeCellList
	UserModelDiff                *client.UserModel                             `json:"user_model_diff"`                   // is actually named _UserModelDiff
}
