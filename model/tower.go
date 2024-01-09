package model

import (
	"elichika/generic"
)

type UserTowerCardUsedCount struct {
	TowerId        int32   `xorm:"pk 'tower_id'" json:"-"`
	CardMasterId   int32   `xorm:"pk 'card_master_id'" json:"card_master_id"`
	UsedCount      int32   `xorm:"'used_count'" json:"used_count"`
	RecoveredCount int32   `xorm:"'recovered_count'" json:"recovered_count"`
	LastUsedAt     int64 `xorm:"'last_used_at'" json:"last_used_at"`
}

type UserTowerVoltageRankingScore struct {
	TowerId int32 `xorm:"pk 'tower_id'"`
	FloorNo int32 `xorm:"pk 'floor_no'"`
	Voltage int32 `xorm:"'voltage'"`
}

type LiveTowerStatus struct {
	TowerId int32 `xorm:"'tower_id'" json:"tower_id"`
	FloorNo int32 `xorm:"'floor_no'" json:"floor_no"`
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
