package model

import (
	"elichika/generic"
)

type UserSteadyVoltageRanking struct {
	SteadyVoltageRankingMasterId int `xorm:"pk 'steady_voltage_ranking_master_id'" json:"steady_voltage_ranking_master_id"`
	SelectedLiveDifficultyId     int `xorm:"'select_live_difficulty_id'" json:"select_live_difficulty_id"`
}

func (usvr *UserSteadyVoltageRanking) Id() int64 {
	return int64(usvr.SteadyVoltageRankingMasterId)
}

func init() {

	TableNameToInterface["u_steady_voltage_ranking"] = generic.UserIdWrapper[UserSteadyVoltageRanking]{}
}
