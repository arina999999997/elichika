package userdata

import (
	"elichika/utils"
)

// TODO(refactor): Move into subsystem
func customBackgroundFinalizer(session *Session) {
	for _, userCustomBackground := range session.UserModel.UserCustomBackgroundById.Map {
		affected, err := session.Db.Table("u_custom_background").Where("user_id = ? AND custom_background_master_id = ?",
			session.UserId, userCustomBackground.CustomBackgroundMasterId).AllCols().Update(userCustomBackground)
		utils.CheckErr(err)
		if affected == 0 {
			GenericDatabaseInsert(session, "u_custom_background", *userCustomBackground)
		}
	}
}
func init() {
	AddFinalizer(customBackgroundFinalizer)
}
