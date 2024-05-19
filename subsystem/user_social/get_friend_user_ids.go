package user_social

import (
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"
)

func GetFriendUserIds(session *userdata.Session) []int32 {
	res := []int32{}
	err := session.Db.Table("u_friend_status").Where("user_id = ? AND request_status = ?",
		session.UserId, enum.FriendRequestStatusFriend).Cols("other_user_id").Find(&res)
	utils.CheckErr(err)
	return res
}
