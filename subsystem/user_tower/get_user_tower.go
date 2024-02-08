package user_tower

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserTower(session *userdata.Session, towerId int32) client.UserTower {
	ptr, exist := session.UserModel.UserTowerByTowerId.Get(towerId)
	if exist {
		return *ptr
	}
	tower := client.UserTower{}
	exist, err := session.Db.Table("u_tower").
		Where("user_id = ? AND tower_id = ?", session.UserId, towerId).Get(&tower)
	utils.CheckErr(err)
	if !exist {
		tower = client.UserTower{
			TowerId:                     towerId,
			ClearedFloor:                0,
			ReadFloor:                   0,
			Voltage:                     0,
			RecoveryPointFullAt:         session.Time.Unix() + 86400,
			RecoveryPointLastConsumedAt: session.Time.Unix(),
		}
	} else {

		// this is to make sure user can always mass recover LP
		// it's the only way to do this since these things are decided by the timestamps alone (+ plus the limit from the database)
		tower.RecoveryPointFullAt = session.Time.Unix() + 86400
		tower.RecoveryPointLastConsumedAt = session.Time.Unix()
	}
	return tower
}
