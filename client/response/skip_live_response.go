package response

import (
	"elichika/client"
)

type SkipLiveResponse struct {
	SkipLiveResult client.SkipLiveResult `json:"skip_live_result"`
	UserModelDiff  *client.UserModel     `json:"user_model_diff"`
}
