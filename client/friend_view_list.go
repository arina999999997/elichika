package client

import (
	"elichika/generic"
)

type FriendViewList struct {
	FriendList                generic.Array[OtherUser] `json:"friend_list"`
	FriendApprovalPendingList generic.Array[OtherUser] `json:"friend_approval_pending_list"`
	FriendApplicationList     generic.Array[OtherUser] `json:"friend_application_list"`
	RecomendPlayerList        generic.Array[OtherUser] `json:"recomend_player_list"`
}
