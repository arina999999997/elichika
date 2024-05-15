package user_social

import (
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
	"elichika/userdata/database"
)

// approve a friend request
// return ok, error key (if not ok)
func ApproveFriendRequestShared(session *userdata.Session, otherUserId int32) (bool, int32) {
	if session.UserId == otherUserId {
		panic("must have different user id")
	}
	// various checks
	if !UserExist(session, otherUserId) {
		return false, enum.FriendFailureTypeApproveNotExist
	}
	if IsMaxFriend(session) {
		return false, enum.FriendFailureTypeApproveMaxFriend
	}
	if IsOtherUserMaxFriend(session, otherUserId) {
		return false, enum.FriendFailureTypeApproveMaxTargetFriend
	}
	// TODO(social): FriendFailureTypeApproveExpired is not implemented as friend requests don't expire for now

	incoming := GetUserFriendStatus(session, otherUserId)
	if !incoming.IsRequestPending {
		return false, enum.FriendFailureTypeApproveNotExist
	}
	incoming.RequestStatus = enum.FriendRequestStatusFriend
	incoming.FriendApprovedAt = generic.NewNullable(session.Time.Unix())
	incoming.IsRequestPending = false
	incoming.IsNew = false

	outgoing := database.UserFriendStatus{
		UserId:           otherUserId,
		OtherUserId:      session.UserId,
		RequestStatus:    enum.FriendRequestStatusFriend,
		FriendApprovedAt: incoming.FriendApprovedAt,
		IsRequestPending: false,
		IsNew:            true,
	}
	UpdateUserFriendStatus(session, incoming)
	UpdateUserFriendStatus(session, outgoing)
	return true, 0
}
