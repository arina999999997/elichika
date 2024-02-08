package user_unlock_scene

import (
	"elichika/client"
	"elichika/userdata"
)

func UnlockScene(session *userdata.Session, unlockSceneType, status int32) {
	// status must be either 1 or 2, any other value and the game will think it doesn't exist at all
	// status = 1 is the initial unlock process, it will show an animation
	// status = 2 is actually unlocked
	userUnlockScene := client.UserUnlockScene{
		UnlockSceneType: unlockSceneType,
		Status:          status,
	}
	session.UserModel.UserUnlockScenesByEnum.Set(unlockSceneType, userUnlockScene)
}