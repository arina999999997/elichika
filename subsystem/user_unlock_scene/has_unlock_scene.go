package user_unlock_scene

import (
	"elichika/userdata"
	"elichika/utils"
)

func HasUnlockScene(session *userdata.Session, unlockSceneType int32) bool {
	exist, err := session.Db.Table("u_unlock_scenes").Where("user_id = ? AND unlock_scene_type = ?",
		session.UserId, unlockSceneType).Exist()
	utils.CheckErr(err)
	return exist
}
