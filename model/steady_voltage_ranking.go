package model

type UserSteadyVoltageRanking struct {
	UserID                       int `xorm:"pk 'user_id'" json:"-"`
	SteadyVoltageRankingMasterID int `xorm:"pk 'steady_voltage_ranking_master_id'" json:"steady_voltage_ranking_master_id"`
	SelectedLiveDifficultyID     int `xorm:"'select_live_difficulty_id'" json:"select_live_difficulty_id"`
}

func (usvr *UserSteadyVoltageRanking) ID() int64 {
	return int64(usvr.SteadyVoltageRankingMasterID)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_steady_voltage_ranking"] = UserSteadyVoltageRanking{}
}
