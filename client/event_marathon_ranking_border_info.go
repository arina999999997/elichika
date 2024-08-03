package client

type EventMarathonRankingBorderInfo struct {
	RankingBorderPoint     int32                               `json:"ranking_border_point"`
	RankingBorderMasterRow EventMarathonRankingBorderMasterRow `json:"ranking_border_master_row"`
}
