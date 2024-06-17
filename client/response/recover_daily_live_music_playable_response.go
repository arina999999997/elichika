package response

import (
	"elichika/client"
)

type RecoverDailyLiveMusicPlayableResponse struct {
	UserModelDiff *client.UserModel `json:"user_model_diff"`
	LiveDaily     client.LiveDaily  `json:"live_daily"`
}
