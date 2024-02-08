package user_tower

import (
	"elichika/userdata"
	"elichika/utils"
)

func userTowerFinalizer(session *userdata.Session) {
	for _, userTower := range session.UserModel.UserTowerByTowerId.Map {
		affected, err := session.Db.Table("u_tower").
			Where("user_id = ? AND tower_id = ?", session.UserId, userTower.TowerId).
			AllCols().Update(*userTower)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_tower", *userTower)
		}
	}
}

func init() {
	userdata.AddFinalizer(userTowerFinalizer)
}
