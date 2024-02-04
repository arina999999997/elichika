package user_custom_background

import (
	"elichika/userdata"
	"elichika/utils"
)

func customBackgroundFinalizer(session *userdata.Session) {
	for _, userCustomBackground := range session.UserModel.UserCustomBackgroundById.Map {
		affected, err := session.Db.Table("u_custom_background").Where("user_id = ? AND custom_background_master_id = ?",
			session.UserId, userCustomBackground.CustomBackgroundMasterId).AllCols().Update(userCustomBackground)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_custom_background", *userCustomBackground)
		}
	}
}

func init() {
	userdata.AddFinalizer(customBackgroundFinalizer)
}
