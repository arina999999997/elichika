package client

import (
	"elichika/generic"
)

type EventMiningRankingBorderMasterRow struct {
	RankingCategory int32                   `json:"ranking_category" enum:"EventRankingCategory"`
	RankingType     int32                   `json:"ranking_type" enum:"EventCommonRankingType"`
	UpperRank       int32                   `json:"upper_rank"`
	LowerRank       generic.Nullable[int32] `json:"lower_rank"`
	DisplayOrder    int32                   `json:"display_order"`
}
