package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchTrainingTreeResponse struct {
	UserCardTrainingTreeCellList generic.List[client.UserCardTrainingTreeCell] `json:"user_card_training_tree_cell_list"` // actually named _UserCardTrainingTreeCellList
}
