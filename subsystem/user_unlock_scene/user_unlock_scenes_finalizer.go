package user_unlock_scene

import (
	"elichika/utils"
	"elichika/userdata"
)

func userUnlockScenesFinalizer(session *userdata.Session) {
	for _, userUnlockScene := range session.UserModel.UserUnlockScenesByEnum.Map {
		affected, err := session.Db.Table("u_unlock_scenes").Where("user_id = ? AND unlock_scene_type = ?",
			session.UserId, userUnlockScene.UnlockSceneType).Update(*userUnlockScene)
		utils.CheckErr(err)
		if affected == 0 { // need to insert
			userdata.GenericDatabaseInsert(session, "u_unlock_scenes", *userUnlockScene)
		}
	}
}
func init() {
	userdata.AddFinalizer(userUnlockScenesFinalizer)
}