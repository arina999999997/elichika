package user_tower

import (
	"elichika/client"
	"elichika/userdata"
)

func UpdateUserTower(session *userdata.Session, tower client.UserTower) {
	session.UserModel.UserTowerByTowerId.Set(tower.TowerId, tower)
}
