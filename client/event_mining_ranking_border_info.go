package client

type EventMiningRankingBorderInfo struct {
	RankingBorderPoint     int32                             `json:"ranking_border_point"`
	RankingBorderMasterRow EventMiningRankingBorderMasterRow `json:"ranking_border_master_row"`
}
