package user_member_guild

import (
	"elichika/client"
	"elichika/subsystem/cache"
	"elichika/userdata"
	"elichika/utils"

	"sort"
)

var (
	getFetchMemberGuildRankingOneTermCache = cache.UniquePointerMap[int32, cache.CachedObject[client.MemberGuildRankingOneTerm]]{}
)

func FetchMemberGuildRankingOneTerm(session *userdata.Session, memberGuildId int32) client.MemberGuildRankingOneTerm {
	cacher := getFetchMemberGuildRankingOneTermCache.Get(-1)
	cacher.Acquire()
	defer cacher.Release()
	if memberGuildId == cacher.Value.MemberGuildId || cacher.ExpireAt <= session.Time.Unix() {
		cacher.ExpireAt = session.Time.Unix() + FetchMemberGuildRankingOneTermCache
		cacher.Value = getFetchMemberGuildRankingOneTermNoCache(session, memberGuildId)
	}
	return *cacher.Value
}

func getFetchMemberGuildRankingOneTermNoCache(session *userdata.Session, memberGuildId int32) *client.MemberGuildRankingOneTerm {
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
	return &ranking
}
