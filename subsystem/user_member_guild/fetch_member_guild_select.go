package user_member_guild

import (
	"elichika/client/response"
	"elichika/userdata"
)

func FetchMemberGuildSelect(session *userdata.Session) response.FetchMemberGuildSelectResponse {
	memberGuildId := GetCurrentMemberGuildId(session) - 1
	if memberGuildId >= 1 {
		return response.FetchMemberGuildSelectResponse{
			PreviousMemberGuildRanking: FetchMemberGuildRankingOneTerm(session, memberGuildId),
		}
	} else {
		return response.FetchMemberGuildSelectResponse{}
	}
}
