package user_tower

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserTowerCardUsedList(session *userdata.Session, towerId int32) generic.List[client.TowerCardUsedCount] {
	list := generic.List[client.TowerCardUsedCount]{}
	err := session.Db.Table("u_tower_card_used_count").
		Where("user_id = ? AND tower_id = ?", session.UserId, towerId).Find(&list.Slice)
	utils.CheckErr(err)
	return list
}
