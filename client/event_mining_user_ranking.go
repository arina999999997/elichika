package client

import (
	"elichika/generic"
)

type EventMiningUserRanking struct {
	PointRankingOrder      int32                   `json:"point_ranking_order"`
	PointRankingTotalPoint generic.Nullable[int32] `json:"point_ranking_total_point"`
	VoltageRankingOrder    int32                   `json:"voltage_ranking_order"`
	VoltageRankingPoint    generic.Nullable[int32] `json:"voltage_ranking_point"`
}
