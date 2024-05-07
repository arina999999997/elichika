package user_member_guild

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"

	"sort"
)

func FetchMemberGuildRankingOneTerm(session *userdata.Session, memberGuildId int32) client.MemberGuildRankingOneTerm {
	// TODO(cache)
	// official server only return 10 of the top ranking unless it's the current one
	// it's likely that this was done using a table for total point and a SELECT ORDER BY LIMIT sql
	// we build the full result for now, the truncation can be done by the caller
	startAt, endAt := GetMemberGuildStartAndEnd(session, memberGuildId)
	ranking := client.MemberGuildRankingOneTerm{
		MemberGuildId: memberGuildId,
		StartAt:       startAt,
		EndAt:         endAt,
	}
	for _, member := range session.Gamedata.Member {
		cell := client.MemberGuildRankingOneTermCell{
			MemberMasterId: member.Id,
		}
		totalPoint, err := session.Db.Table("u_member_guild").
			Where("member_master_id = ? AND member_guild_id = ? AND total_point >= ?",
				member.Id, memberGuildId, session.Gamedata.MemberGuildConstant.JoinConditionPoint).
			OrderBy("total_point DESC").Limit(int(session.Gamedata.MemberGuildConstant.JoinConditionRank)).SumInt(
			&client.UserMemberGuild{}, "total_point")
		utils.CheckErr(err)
		cell.TotalPoint = int32(totalPoint) / session.Gamedata.MemberGuildConstant.JoinConditionRank
		ranking.Channels.Append(cell)
	}
	sort.Slice(ranking.Channels.Slice, func(i, j int) bool {
		return ranking.Channels.Slice[i].TotalPoint > ranking.Channels.Slice[j].TotalPoint
	})
	// the client expect all the number to be present, so we can't handle tieing here
	// this result is not reliable to decide who got what rank
	for i := range ranking.Channels.Slice {
		ranking.Channels.Slice[i].Order = int32(i) + 1
	}
	return ranking
}
