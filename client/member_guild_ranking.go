package client

import (
	"elichika/generic"
)

type MemberGuildRanking struct {
	ViewYear               int32                                   `json:"view_year"`
	NextYear               generic.Nullable[int32]                 `json:"next_year"`
	PreviousYear           generic.Nullable[int32]                 `json:"previous_year"`
	MemberGuildRankingList generic.List[MemberGuildRankingOneTerm] `json:"member_guild_ranking_list"`
}
