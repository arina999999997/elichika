package user_custom_background

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserCustomBackground(session *userdata.Session, customBackgroundMasterId int32) client.UserCustomBackground {
	userCustomBackground := client.UserCustomBackground{}
	exist, err := session.Db.Table("u_custom_background").
		Where("user_id = ? AND custom_background_master_id = ?", session.UserId, customBackgroundMasterId).
		Get(&userCustomBackground)
	utils.CheckErr(err)
	if !exist {
		userCustomBackground = client.UserCustomBackground{
			CustomBackgroundMasterId: customBackgroundMasterId,
			IsNew:                    true,
		}
	}
	return userCustomBackground
}
