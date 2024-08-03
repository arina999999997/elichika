package client

type EventMarathonRankingCell struct {
	Order                    int32                    `json:"order"`
	EventPoint               int32                    `json:"event_point"`
	EventMarathonRankingUser EventMarathonRankingUser `json:"event_marathon_ranking_user"`
}
