package database

import (
	"elichika/generic"
)

// this store the friends and the friend request
type UserFriendStatus struct {
	UserId           int32                   `xorm:"pk 'user_id'"`
	OtherUserId      int32                   `xorm:"pk 'other_user_id'"`
	FriendApprovedAt generic.Nullable[int64] `xorm:"json 'friend_approved_at'"`
	RequestStatus    int32                   `xorm:"'request_status'" enum:"FriendRequestStatus"`
	IsRequestPending bool                    `xorm:"is_request_pending"`
	IsNew            bool                    `xorm:"is_new"`
}

func init() {
	AddTable("u_friend_status", UserFriendStatus{})
}
