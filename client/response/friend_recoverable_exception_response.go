package response

import (
	"elichika/client"
	"elichika/generic"
)

type FriendRecoverableExceptionResponse struct {
	ErrorKey       int32                                   `json:"error_key" enum:"FriendFailureType"`
	FriendViewList generic.Nullable[client.FriendViewList] `json:"friend_view_list"`
}
