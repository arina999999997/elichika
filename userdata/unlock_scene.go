package userdata

import (
	"elichika/model"
	"elichika/utils"
)

// unlock_scene and scene_tips
// unlock_scene unlock the scene, so 1 is for training and so on
// when unlocked, some tips are shown, then scene_tips is used to not show it again
// /sceneTips/saveSceneTipsType

func (session *Session) UnlockScene(unlockSceneType, status int) {
	// status must be either 1 or 2, any other value and the game will think it doesn't exist at all
	// status = 1 is the initial unlock process, it will show an animation
	// status = 2 is actually unlocked
	userUnlockScene := model.UserUnlockScene{
		UserID:          session.UserStatus.UserID,
		UnlockSceneType: unlockSceneType, // not sure what this is
		Status:          status,
	}
	session.UserModel.UserUnlockScenesByEnum.PushBack(userUnlockScene)
}

func unlockSceneFinalizer(session *Session) {
	for _, userUnlockScene := range session.UserModel.UserUnlockScenesByEnum.Objects {
		affected, err := session.Db.Table("u_unlock_scene").Where("user_id = ? AND unlock_scene_type = ?",
			userUnlockScene.UserID, userUnlockScene.UnlockSceneType).Update(userUnlockScene)
		utils.CheckErr(err)
		if affected == 0 { // need to insert
			_, err = session.Db.Table("u_unlock_scene").Insert(userUnlockScene)
			utils.CheckErr(err)
		}
	}
}

func (session *Session) SaveSceneTips(sceneTipsType int) {
	userSceneTips := model.UserSceneTips{
		UserID:        session.UserStatus.UserID,
		SceneTipsType: sceneTipsType,
	}
	session.UserModel.UserSceneTipsByEnum.PushBack(userSceneTips)
}

func sceneTipsFinalizer(session *Session) {
	for _, userSceneTips := range session.UserModel.UserSceneTipsByEnum.Objects {
		_, err := session.Db.Table("u_scene_tips").Insert(userSceneTips)
		utils.CheckErr(err)
	}
}

func init() {
	addFinalizer(unlockSceneFinalizer)
	addFinalizer(sceneTipsFinalizer)
	addGenericTableFieldPopulator("u_unlock_scene", "UserUnlockScenesByEnum")
	addGenericTableFieldPopulator("u_scene_tips", "UserSceneTipsByEnum")
}
