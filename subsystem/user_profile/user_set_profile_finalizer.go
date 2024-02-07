package user_profile

import (
	"elichika/userdata"
	"elichika/utils"
)

// this is only necessary for account importing, because we don't need the delta patch when we change things
func userSetProfileFinalizer(session *userdata.Session) {
	for _, userSetProfile := range session.UserModel.UserSetProfileById.Map {
		affected, err := session.Db.Table("u_set_profile").Where("user_id = ?",
			session.UserId).AllCols().Update(*userSetProfile)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_set_profile", *userSetProfile)
		}
	}
}

func init() {
	userdata.AddFinalizer(userSetProfileFinalizer)
}
