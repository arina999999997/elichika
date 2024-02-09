package user_scene_tips

import (
	"elichika/userdata"
)

func userSceneTipsFinalizer(session *userdata.Session) {
	for _, userSceneTips := range session.UserModel.UserSceneTipsByEnum.Map {
		userdata.GenericDatabaseInsert(session, "u_scene_tips", *userSceneTips)
	}
}

func init() {
	userdata.AddFinalizer(userSceneTipsFinalizer)
}
