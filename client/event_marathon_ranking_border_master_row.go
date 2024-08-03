package client

import (
	"elichika/generic"
)

type EventMarathonRankingBorderMasterRow struct {
	RankingType  int32                   `json:"ranking_type" enum:"EventCommonRankingType"`
	UpperRank    int32                   `json:"upper_rank"`
	LowerRank    generic.Nullable[int32] `json:"lower_rank"`
	DisplayOrder int32                   `json:"display_order"`
}
