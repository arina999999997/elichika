package user_tower

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func UpdateUserTowerCardUsed(session *userdata.Session, towerId int32, card client.TowerCardUsedCount) {
	affected, err := session.Db.Table("u_tower_card_used_count").
		Where("user_id = ? AND tower_id = ? AND card_master_id = ?", session.UserId, towerId, card.CardMasterId).
		AllCols().Update(card)
	utils.CheckErr(err)
	if affected == 0 {
		type Wrapper struct {
			Card    client.TowerCardUsedCount `xorm:"extends"`
			TowerId int32                     `xorm:"pk 'tower_id'"`
		}
		userdata.GenericDatabaseInsert(session, "u_tower_card_used_count", Wrapper{
			Card:    card,
			TowerId: towerId,
		})
	}
}
