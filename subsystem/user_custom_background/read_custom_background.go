package user_custom_background

import (
	"elichika/userdata"
)

// the new marker really handled, there is no new marker in game, and there is no request to clear it
// the server seems to mark them as used when they are set, probably only to collect data on what people like
func ReadCustomBackground(session *userdata.Session, backgroundId int32) {
	userCustomBackground := GetUserCustomBackground(session, backgroundId)
	if userCustomBackground.IsNew {
		userCustomBackground.IsNew = false
		UpdateUserCustomBackground(session, userCustomBackground)
	}
}