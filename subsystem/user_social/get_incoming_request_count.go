package user_social

import (
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"
)

func GetIncomingRequestCount(session *userdata.Session, userId int32) int32 {
	count, err := session.Db.Table("u_friend_status").Where("user_id = ? AND request_status = ?", userId,
		enum.FriendRequestStatusNone).Count()
	utils.CheckErr(err)
	return int32(count)
}
