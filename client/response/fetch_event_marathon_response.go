package response

import (
	"elichika/client"
)

type FetchEventMarathonResponse struct {
	EventMarathonTopStatus client.EventMarathonTopStatus `json:"event_marathon_top_status"`
	UserModelDiff          *client.UserModel             `json:"user_model_diff"`
}
