package response

import (
	"elichika/client"
)

type FinishLiveResponse struct {
	LiveResult    client.LiveResult `json:"live_result"`
	UserModelDiff *client.UserModel `json:"user_model_diff"`
}
