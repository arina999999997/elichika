package client

type VoltageRankingCell struct {
	Order              int32              `json:"order"`
	VoltagePoint       int32              `json:"voltage_point"`
	VoltageRankingUser VoltageRankingUser `json:"voltage_ranking_user"`
}
