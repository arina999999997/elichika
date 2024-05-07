package user_member_guild

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"

	"sort"
)

func FetchMemberGuildRankingAnimationInfos(session *userdata.Session, memberGuildId int32) generic.List[client.MemberGuildRankingAnimationInfo] {
	list := generic.List[client.MemberGuildRankingAnimationInfo]{}
	for _, member := range session.Gamedata.Member {
		totalPoint, err := session.Db.Table("u_member_guild").
			Where("member_master_id = ? AND member_guild_id = ? AND total_point >= ?",
				member.Id, memberGuildId, session.Gamedata.MemberGuildConstant.JoinConditionPoint).
			OrderBy("total_point DESC").Limit(int(session.Gamedata.MemberGuildConstant.JoinConditionRank)).SumInt(
			&client.UserMemberGuild{}, "total_point")
		utils.CheckErr(err)
		list.Append(client.MemberGuildRankingAnimationInfo{
			MemberMasterId:          member.Id,
			MemberGuildRankingPoint: int32(totalPoint),
		})
	}
	sort.Slice(list.Slice, func(i, j int) bool {
		return list.Slice[i].MemberGuildRankingPoint > list.Slice[j].MemberGuildRankingPoint
	})
	for i := range list.Slice {
		list.Slice[i].MemberGuildRankingOrder = int32(i + 1)
	}
	return list
}

// "bias" the ranking toward a member by putting them ahead of their equal range
// this is actually important to deliver the correct the reward without too much hassle
func biasRanking(list *generic.List[client.MemberGuildRankingAnimationInfo], memberMasterId int32) {
	for i, info := range list.Slice {
		if info.MemberMasterId == memberMasterId {
			for j, otherInfo := range list.Slice {
				if otherInfo.MemberGuildRankingPoint == info.MemberGuildRankingPoint {
					for k := i; k > j; k-- {
						list.Slice[k].MemberMasterId = list.Slice[k-1].MemberMasterId
					}
					list.Slice[j].MemberMasterId = memberMasterId
					return
				}
			}
		}
	}
}
