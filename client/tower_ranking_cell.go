package client

type TowerRankingCell struct {
	Order            int32            `json:"order"`
	SumVoltage       int32            `json:"sum_voltage"`
	TowerRankingUser TowerRankingUser `json:"tower_ranking_user"`
}
