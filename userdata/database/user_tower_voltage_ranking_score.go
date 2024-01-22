package database

import (
	"elichika/generic"
)

type UserTowerVoltageRankingScore struct {
	TowerId int32 `xorm:"pk 'tower_id'"`
	FloorNo int32 `xorm:"pk 'floor_no'"`
	Voltage int32 `xorm:"'voltage'"`
}

func init() {
	AddTable("u_tower_voltage_ranking_score", generic.UserIdWrapper[UserTowerVoltageRankingScore]{})
}
