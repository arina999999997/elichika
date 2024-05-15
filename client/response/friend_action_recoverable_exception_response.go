package response

import (
	"elichika/client"
	"elichika/generic"
)

type FriendActionRecoverableExceptionResponse struct {
	ErrorKey     int32                              `json:"error_key" enum:"FriendFailureType"`
	TargetPlayer generic.Nullable[client.OtherUser] `json:"target_player"`
}
