package user_member_guild

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
)

// fetch the member guild ranking of a year
// all of the member guild are fetched, including past one and current one (if any)
// the cutoff is based on the end date instead of the beginning date
// for now we just loop through the member guild id, maybe we can revisit and make it better if we want this it work after 2038
func FetchMemberGuildRankingYear(session *userdata.Session, viewYear int32) client.MemberGuildRanking {
	lastId := GetCurrentMemberGuildId(session)
	if !IsMemberGuildRankingPeriod(session) {
		lastId--
	}
	if viewYear == 0 {
		viewYear = GetMemberGuildIdYear(session, lastId)
	}

	memberGuildRanking := client.MemberGuildRanking{
		ViewYear: viewYear,
	}
	for id := int32(1); id <= lastId; id++ {
		year := GetMemberGuildIdYear(session, id)
		if year < viewYear {
			memberGuildRanking.PreviousYear = generic.NewNullable(year)
			continue
		} else if year == viewYear {
			ranking := FetchMemberGuildRankingOneTerm(session, id)
			if (id != lastId) && (ranking.Channels.Size() > 10) { // if not last then truncate to 10 channels
				ranking.Channels.Slice = ranking.Channels.Slice[0:10]
			}
			memberGuildRanking.MemberGuildRankingList.Append(ranking)
		} else {
			memberGuildRanking.NextYear = generic.NewNullable(year)
			break
		}
	}
	return memberGuildRanking
}
