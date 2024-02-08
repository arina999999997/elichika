package user_scene_tips

import (
	"elichika/client"
	"elichika/userdata"
)

func SaveUserSceneTips(session *userdata.Session, sceneTipsType int32) {
	userSceneTips := client.UserSceneTips{
		SceneTipsType: sceneTipsType,
	}
	session.UserModel.UserSceneTipsByEnum.Set(sceneTipsType, userSceneTips)
}