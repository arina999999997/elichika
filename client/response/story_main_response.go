package response

import (
	"elichika/client"
	"elichika/generic"
)

type StoryMainResponse struct {
	UserModelDiff    *client.UserModel             `json:"user_model_diff"`
	FirstClearReward generic.Array[client.Content] `json:"first_clear_reward"`
}
