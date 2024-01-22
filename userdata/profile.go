package userdata

import (
	"elichika/client"
	"elichika/utils"
)

func FetchDBProfile(userId int32, result interface{}) {
	exist, err := Engine.Table("u_status").Where("user_id = ?", userId).Get(result)
	utils.CheckErrMustExist(err, exist)
}

// fetch profile of another user, from session.UserId's perspective
// it's possible that otherUserId == session.UserId

func (session *Session) GetOtherUserSetProfile(otherUserId int32) client.UserSetProfile {
	p := client.UserSetProfile{}
	_, err := session.Db.Table("u_set_profile").Where("user_id = ?", otherUserId).Get(&p)
	utils.CheckErr(err)
	return p
}

func (session *Session) GetUserSetProfile() client.UserSetProfile {
	return session.GetOtherUserSetProfile(session.UserId)
}

// doesn't need to return delta patch or submit at the start because we would need to fetch profile everytime we need this thing
func (session *Session) SetUserSetProfile(userSetProfile client.UserSetProfile) {
	affected, err := session.Db.Table("u_set_profile").Where("user_id = ?", session.UserId).
		AllCols().Update(&userSetProfile)
	utils.CheckErr(err)
	if affected == 0 {
		// need to insert
		genericDatabaseInsert(session, "u_set_profile", userSetProfile)
	}
}

func userSetProfileFinalizer(session *Session) {
	for _, userSetProfile := range session.UserModel.UserSetProfileById.Map {
		affected, err := session.Db.Table("u_set_profile").Where("user_id = ?",
			session.UserId).AllCols().Update(*userSetProfile)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_set_profile", *userSetProfile)
		}
	}
}

func init() {
	addFinalizer(userSetProfileFinalizer)
}
