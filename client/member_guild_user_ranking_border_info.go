package client

import (
	"elichika/generic"
)

type MemberGuildUserRankingBorderInfo struct {
	RankingBorderPoint int32                   `json:"ranking_border_point"`
	UpperRank          int32                   `json:"upper_rank"`
	LowerRank          generic.Nullable[int32] `json:"lower_rank"`
	DisplayOrder       int32                   `json:"display_order"`
}
