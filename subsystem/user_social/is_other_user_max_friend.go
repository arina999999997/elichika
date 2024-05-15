package user_social

import (
	"elichika/userdata"
)

func IsOtherUserMaxFriend(session *userdata.Session, otherUserId int32) bool {
	friendCount := GetFriendCount(session, otherUserId)
	otherUserStatus := GetOtherUserStatus(session, otherUserId)
	return friendCount >= session.Gamedata.UserRank[otherUserStatus.Rank].FriendLimit
}
