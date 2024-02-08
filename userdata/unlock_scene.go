package userdata

import (
	"elichika/client"
	"elichika/utils"
)

// TODO(refactor): Move into subsystem
// unlock_scene and scene_tips
// unlock_scene unlock the scene, so 1 is for training and so on
// when unlocked, some tips are shown, then scene_tips is used to not show it again
// /sceneTips/saveSceneTipsType

func (session *Session) UnlockScene(unlockSceneType, status int32) {
	// status must be either 1 or 2, any other value and the game will think it doesn't exist at all
	// status = 1 is the initial unlock process, it will show an animation
	// status = 2 is actually unlocked
	userUnlockScene := client.UserUnlockScene{
		UnlockSceneType: unlockSceneType,
		Status:          status,
	}
	session.UserModel.UserUnlockScenesByEnum.Set(unlockSceneType, userUnlockScene)
}

func unlockSceneFinalizer(session *Session) {
	for _, userUnlockScene := range session.UserModel.UserUnlockScenesByEnum.Map {
		affected, err := session.Db.Table("u_unlock_scenes").Where("user_id = ? AND unlock_scene_type = ?",
			session.UserId, userUnlockScene.UnlockSceneType).Update(*userUnlockScene)
		utils.CheckErr(err)
		if affected == 0 { // need to insert
			GenericDatabaseInsert(session, "u_unlock_scenes", *userUnlockScene)
		}
	}
}

func (session *Session) SaveSceneTips(sceneTipsType int32) {
	userSceneTips := client.UserSceneTips{
		SceneTipsType: sceneTipsType,
	}
	session.UserModel.UserSceneTipsByEnum.Set(sceneTipsType, userSceneTips)
}

func sceneTipsFinalizer(session *Session) {
	for _, userSceneTips := range session.UserModel.UserSceneTipsByEnum.Map {
		GenericDatabaseInsert(session, "u_scene_tips", *userSceneTips)
	}
}

func init() {
	AddFinalizer(unlockSceneFinalizer)
	AddFinalizer(sceneTipsFinalizer)
}
