package user_social

import (
	"elichika/enum"
	"elichika/log"
	"elichika/userdata"
)

// cancel a friend request
// return ok, error key (if not ok)
func CancelFriendRequestShared(session *userdata.Session, otherUserId int32) (bool, int32) {
	if session.UserId == otherUserId {
		log.Panic("must have different user id")
	}

	if IsFriend(session, otherUserId) { // other player accepted before this player cancel
		return false, enum.FriendFailureTypeCancelAlreadyFriend
	}

	RemoveConnection(session, otherUserId)
	return true, 0
}
