package response

import (
	"elichika/client"
)

type FetchEventMiningResponse struct {
	EventMiningTopStatus client.EventMiningTopStatus `json:"event_mining_top_status"`
	UserModelDiff        *client.UserModel           `json:"user_model_diff"`
}
