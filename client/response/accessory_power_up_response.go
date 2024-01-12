package response

import (
	"elichika/client"
)

type AccessoryPowerUpResponse struct {
	Success       int32             `json:"success" enum:"AccessoryLevelUpSuccess"`
	DoPowerUp     client.DoPowerUp  `json:"do_power_up"`
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
