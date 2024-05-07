package user_member_guild

import (
	"elichika/client/response"
	"elichika/userdata"
)

func FetchMemberGuildRanking(session *userdata.Session) response.FetchMemberGuildRankingResponse {
	// ranking between channels is the same as fetching per year, in fact the whole year is returned and cached on client side
	resp := response.FetchMemberGuildRankingResponse{
		MemberGuildRanking: FetchMemberGuildRankingYear(session, 0),
	}
	// ranking within channels has current ranking and previous ranking
	id := GetCurrentMemberGuildId(session)
	resp.MemberGuildUserRankingList.Append(FetchMemberGuildUserRanking(session, id))
	if id > 1 {
		previousUserMemberGuild := GetUserMemberGuild(session, id-1)
		if previousUserMemberGuild.TotalPoint > 0 {
			resp.MemberGuildUserRankingList.Append(FetchMemberGuildUserRanking(session, id-1))
		}
	}
	return resp
}
