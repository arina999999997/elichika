package user_social

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUserStatus(session *userdata.Session, otherUserId int32) client.UserStatus {
	if session.UserId == otherUserId {
		return *session.UserStatus
	}
	otherUserStatus := client.UserStatus{}
	exist, err := session.Db.Table("u_status").Where("user_id = ?", otherUserId).Get(&otherUserStatus)
	utils.CheckErrMustExist(err, exist)
	return otherUserStatus
}
