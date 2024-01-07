package model

type UserSteadyVoltageRanking struct {
	UserId                       int `xorm:"pk 'user_id'" json:"-"`
	SteadyVoltageRankingMasterId int `xorm:"pk 'steady_voltage_ranking_master_id'" json:"steady_voltage_ranking_master_id"`
	SelectedLiveDifficultyId     int `xorm:"'select_live_difficulty_id'" json:"select_live_difficulty_id"`
}

func (usvr *UserSteadyVoltageRanking) Id() int64 {
	return int64(usvr.SteadyVoltageRankingMasterId)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_steady_voltage_ranking"] = UserSteadyVoltageRanking{}
}
