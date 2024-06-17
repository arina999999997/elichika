package user_unlock_scene

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func UnlockScene(session *userdata.Session, unlockSceneType, status int32) {
	// status must be either 1 or 2, any other value and the game will think it doesn't exist at all
	// status = 1 is the initial unlock process, it will show an animation
	// status = 2 is actually unlocked
	userUnlockScene := client.UserUnlockScene{}

	_, err := session.Db.Table("u_unlock_scenes").Where("user_id = ? AND unlock_scene_type = ?",
		session.UserId, unlockSceneType).Get(&userUnlockScene)
	utils.CheckErr(err)
	if userUnlockScene.Status >= status {
		// already unlocked
		return
	}
	session.UserModel.UserUnlockScenesByEnum.Set(unlockSceneType, client.UserUnlockScene{
		UnlockSceneType: unlockSceneType,
		Status:          status,
	})
}
