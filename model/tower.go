package model

import (
	"elichika/client"
	"elichika/generic"
)

type UserTowerCardUsedCount struct {
	TowerId int32 `xorm:"pk 'tower_id'" json:"-"`
	client.TowerCardUsedCount
}

type UserTowerVoltageRankingScore struct {
	TowerId int32 `xorm:"pk 'tower_id'"`
	FloorNo int32 `xorm:"pk 'floor_no'"`
	Voltage int32 `xorm:"'voltage'"`
}

func init() {
	TableNameToInterface["u_tower_card_used_count"] = generic.UserIdWrapper[UserTowerCardUsedCount]{}
	TableNameToInterface["u_tower_voltage_ranking_score"] = generic.UserIdWrapper[UserTowerVoltageRankingScore]{}
}
