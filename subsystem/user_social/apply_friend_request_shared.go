package user_social

import (
	"elichika/enum"
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

// add a friend request
// special behavior:
// - If the other user already have a request directed at this user, sending them a request will actually approve the existing request
// return ok, error key (if not ok)
func ApplyFriendRequestShared(session *userdata.Session, otherUserId int32) (bool, int32) {
	if session.UserId == otherUserId {
		panic("must have different user id")
	}
	// various checks
	if !UserExist(session, otherUserId) {
		return false, enum.FriendFailureTypeSearchNotExist
	}
	if IsMaxFriend(session) {
		return false, enum.FriendFailureTypeRequestMaxFriend
	}
	if IsOtherUserMaxFriend(session, otherUserId) {
		return false, enum.FriendFailureTypeRequestMaxTargetFriend
	}
	if GetOutgoingRequestCount(session) >= session.Gamedata.ConstantInt[enum.ConstantIntFriendApplicationMaxSendCount] {
		return false, enum.FriendFailureTypeRequestMaxApplication
	}
	if GetIncomingRequestCount(session, otherUserId) >= session.Gamedata.ConstantInt[enum.ConstantIntFriendApplicationMaxReceiveCount] {
		return false, enum.FriendFailureTypeRequestMaxTargetApplication
	}
	if IsFriend(session, otherUserId) {
		return false, enum.FriendFailureTypeRequestAlreadyFriend
	}

	outgoing := GetUserFriendStatus(session, otherUserId)
	if outgoing.IsRequestPending { // other user sent a request
		good, errorKey := ApproveFriendRequestShared(session, otherUserId)
		return good, errorKey
	}

	// no existing request, initiate one
	// successful checks, add the friend request
	outgoing = database.UserFriendStatus{
		UserId:           session.UserId,
		OtherUserId:      otherUserId,
		RequestStatus:    enum.FriendRequestStatusRequest,
		IsRequestPending: false,
		IsNew:            false,
	}
	incoming := database.UserFriendStatus{
		UserId:           otherUserId,
		OtherUserId:      session.UserId,
		RequestStatus:    enum.FriendRequestStatusNone,
		IsRequestPending: true,
		IsNew:            true,
	}
	_, err := session.Db.Table("u_friend_status").Insert(incoming)
	utils.CheckErr(err)
	_, err = session.Db.Table("u_friend_status").Insert(outgoing)
	utils.CheckErr(err)
	return true, 0
}
