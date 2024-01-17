package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchMemberGuildRankingResponse struct {
	MemberGuildRanking         client.MemberGuildRanking                   `json:"member_guild_ranking"`
	MemberGuildUserRankingList generic.List[client.MemberGuildUserRanking] `json:"member_guild_user_ranking_list"`
}
