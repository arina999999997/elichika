package userdata

import (
	"elichika/utils"
)

func towerFinalizer(session *Session) {
	for _, userTower := range session.UserModel.UserTowerByTowerID.Objects {
		affected, err := session.Db.Table("u_tower").
			Where("user_id = ? AND tower_id = ?",
				session.UserStatus.UserID, userTower.TowerID).
			AllCols().Update(userTower)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_tower").
				Insert(userTower)
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(towerFinalizer)
	addGenericTableFieldPopulator("u_tower", "UserTowerByTowerID")
}
