package user_social

import (
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"
)

func GetOutgoingRequestCount(session *userdata.Session) int32 {
	count, err := session.Db.Table("u_friend_status").Where("user_id = ? AND request_status = ?", session.UserId,
		enum.FriendRequestStatusRequest).Count()
	utils.CheckErr(err)
	return int32(count)
}
