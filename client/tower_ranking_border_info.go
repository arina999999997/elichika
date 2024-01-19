package client

type TowerRankingBorderInfo struct {
	RankingBorderVoltage   int32                       `json:"ranking_border_voltage"`
	RankingBorderMasterRow TowerRankingBorderMasterRow `json:"ranking_border_master_row"`
}
