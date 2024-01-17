package client

import (
	"elichika/generic"
)

type MemberGuildUserRanking struct {
	MemberGuildId  int32                                          `json:"member_guild_id"`
	TopRanking     generic.List[MemberGuildUserRankingCell]       `json:"top_ranking"`
	MyRanking      generic.List[MemberGuildUserRankingCell]       `json:"my_ranking"`
	RankingBorders generic.List[MemberGuildUserRankingBorderInfo] `json:"ranking_borders"`
}
