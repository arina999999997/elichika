package response

import (
	"elichika/client"
	"elichika/generic"
)

type FriendActionResponse struct {
	SuccessType  int32                              `json:"success_type" enum:"FriendSuccessType"`
	TargetPlayer generic.Nullable[client.OtherUser] `json:"target_player"`
}
