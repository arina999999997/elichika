package user_member_guild

import (
	"elichika/client"
	"elichika/generic"
	"elichika/subsystem/user_social"
	"elichika/userdata"
	"elichika/utils"
)

func FetchMemberGuildUserRanking(session *userdata.Session, memberGuildId int32) client.MemberGuildUserRanking {
	ranking := client.MemberGuildUserRanking{
		MemberGuildId: memberGuildId,
	}
	// top ranking returned is the top 100 for GL, so it could be the 3rd ranking bracket, or it could be top 100 in jp as well
	// for simplicity this is hard coded to 100
	// top ranking always use the current user id
	memberMasterId := session.UserStatus.MemberGuildMemberMasterId.Value
	userMemberGuilds := []generic.UserIdWrapper[client.UserMemberGuild]{}
	err := session.Db.Table("u_member_guild").
		Where("member_master_id = ? AND member_guild_id = ? AND total_point >= ?",
			memberMasterId, memberGuildId, session.Gamedata.MemberGuildConstant.JoinConditionPoint).
		OrderBy("total_point DESC").Limit(100).Find(&userMemberGuilds)
	utils.CheckErr(err)
	for i, user := range userMemberGuilds {
		memberGuild := user.Object
		ranking.TopRanking.Append(client.MemberGuildUserRankingCell{
			TotalPoint:                     memberGuild.TotalPoint,
			MemberGuildUserRankingUserData: user_social.GetMemberGuildUserRankingUserData(session, user.UserId),
		})
		if (i == 0) || (ranking.TopRanking.Slice[i].TotalPoint != ranking.TopRanking.Slice[i-1].TotalPoint) {
			ranking.TopRanking.Slice[i].Order = generic.NewNullable(int32(i + 1))
		} else {
			ranking.TopRanking.Slice[i].Order = ranking.TopRanking.Slice[i-1].Order
		}
	}

	// own ranking
	userMemberGuild := GetUserMemberGuild(session, memberGuildId)
	userMemberGuilds = []generic.UserIdWrapper[client.UserMemberGuild]{}
	if (userMemberGuild.MemberMasterId == memberMasterId) && (userMemberGuild.TotalPoint > 0) {
		// only fetch ranking if user was in the same channel as currently, not sure if this is correct but it makes more sense
		// the limit for this seems to be [last number that mod 20 = 1, current user's rank + 40]
		// so rank 1 load the range [1, 41], rank 20 load the range [1, 60] and, rank 22 load the range [21, 62]
		// the loading down 40 people is observed in GL network record, not sure about the loading up or JP
		// there can also be a problem if there are more than a 40 ways tie(?), but let's not bother with that case
		rank, err := session.Db.Table("u_member_guild").Where("member_master_id = ? AND member_guild_id = ? AND total_point > ?",
			userMemberGuild.MemberMasterId, memberGuildId, userMemberGuild.TotalPoint).OrderBy("total_point DESC").
			Limit(int(session.Gamedata.MemberGuildRankingRewardInside[userMemberGuild.MemberMasterId].RankNumberLimit)).Count()
		utils.CheckErr(err)
		rank++
		if int32(rank) > session.Gamedata.MemberGuildRankingRewardInside[userMemberGuild.MemberMasterId].RankNumberLimit {
			rank = 0
		}
		lowerRank := rank + 40
		if rank%20 != 1 {
			rank -= rank % 20
			rank++
		}

		err = session.Db.Table("u_member_guild").
			Where("member_master_id = ? AND member_guild_id = ? AND total_point >= ?",
				userMemberGuild.MemberMasterId, memberGuildId, session.Gamedata.MemberGuildConstant.JoinConditionPoint).
			OrderBy("total_point DESC").Limit(int(lowerRank-rank+1), int(rank-1)).Find(&userMemberGuilds)
		utils.CheckErr(err)

		for i, user := range userMemberGuilds {
			memberGuild := user.Object
			ranking.MyRanking.Append(client.MemberGuildUserRankingCell{
				TotalPoint:                     memberGuild.TotalPoint,
				MemberGuildUserRankingUserData: user_social.GetMemberGuildUserRankingUserData(session, user.UserId),
			})
			ranking.MyRanking.Slice[i].Order = generic.NewNullable(int32(i) + int32(rank))
		}
	}

	// ranking border
	for i, step := range session.Gamedata.MemberGuildRankingRewardInside[memberMasterId].Steps {
		border := client.MemberGuildUserRankingBorderInfo{
			UpperRank:    step.UpperRank,
			DisplayOrder: int32(i),
		}
		if step.LowerRank != nil {
			border.LowerRank = generic.NewNullable(*step.LowerRank)
			exist, err := session.Db.Table("u_member_guild").
				Where("member_master_id = ? AND member_guild_id = ? AND total_point >= ?",
					memberMasterId, memberGuildId, session.Gamedata.MemberGuildConstant.JoinConditionPoint).
				OrderBy("total_point DESC").Limit(1, int(*step.LowerRank)-1).Cols("total_point").Get(&border.RankingBorderPoint)
			utils.CheckErr(err)
			if !exist {
				border.RankingBorderPoint = 0
			}
		}
		ranking.RankingBorders.Append(border)
		if border.RankingBorderPoint == 0 { // skip further border
			break
		}
	}
	return ranking
}
