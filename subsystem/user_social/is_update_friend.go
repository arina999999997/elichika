package user_social

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func IsUpdateFriend(session *userdata.Session) bool {
	exist, err := session.Db.Table("u_friend_status").Where("user_id = ? AND is_new != 0", session.UserId).Exist(&database.UserFriendStatus{})
	utils.CheckErr(err)
	return exist
}
