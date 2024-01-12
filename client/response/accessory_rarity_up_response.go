package response

import (
	"elichika/client"
)

type AccessoryRarityUpResponse struct {
	DoRarityUp    client.DoRarityUp `json:"do_rarity_up"`
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
