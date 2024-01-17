package client

type LoveRankingData struct {
	RankingUser RankingUser `json:"ranking_user"`
	Order       int32       `json:"order"`
	LovePoint   int32       `json:"love_point"`
}
