package client

type EventMiningRankingCell struct {
	Order                  int32                  `json:"order"`
	EventPoint             int32                  `json:"event_point"`
	EventMiningRankingUser EventMiningRankingUser `json:"event_mining_ranking_user"`
}
