package user_custom_background

import (
	"elichika/client"
	"elichika/userdata"
)

func UpdateUserCustomBackground(session *userdata.Session, userCustomBackground client.UserCustomBackground) {
	session.UserModel.UserCustomBackgroundById.Set(userCustomBackground.CustomBackgroundMasterId, userCustomBackground)
}
