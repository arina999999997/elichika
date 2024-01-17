package response

import (
	"elichika/client"
)

type FetchMemberGuildSelectResponse struct {
	PreviousMemberGuildRanking client.MemberGuildRankingOneTerm `json:"previous_member_guild_ranking"`
}
