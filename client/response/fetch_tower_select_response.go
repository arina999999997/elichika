package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchTowerSelectResponse struct {
	TowerIds      generic.Array[int32] `json:"tower_ids"`
	UserModelDiff *client.UserModel    `json:"user_model_diff"`
}
