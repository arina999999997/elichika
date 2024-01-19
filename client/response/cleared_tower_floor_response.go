package response

import (
	"elichika/client"
)

type ClearedTowerFloorResponse struct {
	IsShowUnlockEffect bool              `json:"is_show_unlock_effect"`
	UserModelDiff      *client.UserModel `json:"user_model_diff"`
}
