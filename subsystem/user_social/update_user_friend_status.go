package user_social

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func UpdateUserFriendStatus(session *userdata.Session, status database.UserFriendStatus) {
	// update only, DO NOT not insert
	affected, err := session.Db.Table("u_friend_status").Where("user_id = ? AND other_user_id = ?",
		status.UserId, status.OtherUserId).AllCols().Update(&status)
	utils.CheckErr(err)
	if affected == 0 {
		panic("friend status doesn't exist")
	}
}
