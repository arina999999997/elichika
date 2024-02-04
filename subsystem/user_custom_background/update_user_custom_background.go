package user_custom_background

import (
	"elichika/userdata"
	"elichika/client"
)

func UpdateUserCustomBackground(session *userdata.Session, userCustomBackground client.UserCustomBackground) {
	session.UserModel.UserCustomBackgroundById.Set(userCustomBackground.CustomBackgroundMasterId, userCustomBackground)
}