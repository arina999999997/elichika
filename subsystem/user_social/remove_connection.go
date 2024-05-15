package user_social

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func RemoveConnection(session *userdata.Session, otherUserId int32) {
	// delete the connection with the other player no matter what
	_, err := session.Db.Table("u_friend_status").Where("user_id = ? AND other_user_id = ?", session.UserId, otherUserId).
		Delete(&database.UserFriendStatus{})
	utils.CheckErr(err)
	_, err = session.Db.Table("u_friend_status").Where("user_id = ? AND other_user_id = ?", otherUserId, session.UserId).
		Delete(&database.UserFriendStatus{})
	utils.CheckErr(err)
}
