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

type TowerLive struct {
	TowerId       *int32 `xorm:"tower_id" json:"tower_id"`
	FloorNo       *int32 `xorm:"tower_floor_no" json:"floor_no"`
	TargetVoltage *int32 `xorm:"tower_target_voltage" json:"target_voltage"`
	StartVoltage  *int32 `xorm:"tower_start_voltage" json:"start_voltage"`
}

func init() {

	TableNameToInterface["u_tower_card_used"] = generic.UserIdWrapper[UserTowerCardUsedCount]{}
	TableNameToInterface["u_tower_voltage_ranking_score"] = generic.UserIdWrapper[UserTowerVoltageRankingScore]{}
}
