package response

import (
	"elichika/client"
)

type FetchMemberGuildRankingYearResponse struct {
	MemberGuildRanking client.MemberGuildRanking `json:"member_guild_ranking"`
}
