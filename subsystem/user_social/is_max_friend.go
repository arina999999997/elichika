package user_social

import (
	"elichika/userdata"
)

func IsMaxFriend(session *userdata.Session) bool {
	friendCount := GetFriendCount(session, session.UserId)
	return friendCount >= session.Gamedata.UserRank[session.UserModel.UserStatus.Rank].FriendLimit
}
