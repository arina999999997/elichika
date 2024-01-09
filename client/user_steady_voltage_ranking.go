package client

type UserSteadyVoltageRanking struct {
	SteadyVoltageRankingMasterId int32 `xorm:"pk 'steady_voltage_ranking_master_id'" json:"steady_voltage_ranking_master_id"`
	SelectedLiveDifficultyId     int32 `xorm:"'select_live_difficulty_id'" json:"select_live_difficulty_id"`
}

func (usvr *UserSteadyVoltageRanking) Id() int64 {
	return int64(usvr.SteadyVoltageRankingMasterId)
}
