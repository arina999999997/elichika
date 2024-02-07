package user_profile

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetOtherUserSetProfile(session *userdata.Session, otherUserId int32) client.UserSetProfile {
	p := client.UserSetProfile{}
	_, err := session.Db.Table("u_set_profile").Where("user_id = ?", otherUserId).Get(&p)
	utils.CheckErr(err)
	return p
}
