package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchLoveRankingResponse struct {
	LoveRankingData generic.Array[client.LoveRankingData] `json:"love_ranking_data"`
	MyRankingOrder  generic.Nullable[int32]               `json:"my_ranking_order"`
}
