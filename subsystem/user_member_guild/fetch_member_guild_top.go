package user_member_guild

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_info_trigger"
	"elichika/subsystem/user_present"
	"elichika/userdata"
	"elichika/utils"

	"fmt"
)

// this response has 3 different states it can return:
// - Display the result of the previous member guild and display the reward
//   - Once after the previous guild end, checked using a trigger created by bootstrap and stored in database
//
// - Display the result of the previous member guild only
//   - When IsBeforeMemberGuildRankingPeriod return true
//
// - Display the current member guild and current ranking
//   - If the above doesn't matches
func FetchMemberGuildTop(session *userdata.Session) response.FetchMemberGuildTopResponse {
	resp := response.FetchMemberGuildTopResponse{
		UserModelDiff: &session.UserModel,
	}
	currentMemberGuildId := GetCurrentMemberGuildId(session)
	previousMemberGuildId := currentMemberGuildId - 1
	if previousMemberGuildId > 0 {
		resp.MemberGuildTopStatus.MemberGuildRankingResultAnimationInfo =
			FetchMemberGuildRankingAnimationInfos(session, previousMemberGuildId)
	}

	resultTriggers := []client.UserInfoTriggerBasic{}
	err := session.Db.Table("u_info_trigger_basic").Where("user_id = ? AND info_trigger_type = ?",
		session.UserId, enum.InfoTriggerTypeMemberGuildRankingShowResult).Find(&resultTriggers)
	utils.CheckErr(err)
	hasResultTrigger := false
	for i, trigger := range resultTriggers {
		if trigger.LimitAt.Value >= session.Time.Unix() {
			hasResultTrigger = true
		}
		if i == 0 {
			// always remove the triggers of this type, if any
			user_info_trigger.DeleteTriggerBasicByType(session, enum.InfoTriggerTypeMemberGuildRankingShowResult)
		}
	}
	if hasResultTrigger ||
		(IsBeforeMemberGuildRankingPeriod(session) && session.UserStatus.MemberGuildMemberMasterId.HasValue) {
		// result display and reward awarding
		// no need to add the voice, it will be added by a /navi/saveUserNaviVoice call later on
		// TODO(extra) technically this lead to the risk of the voice not being unlocked
		resp.MemberGuildTopStatus.IsTopRankingDisplay = hasResultTrigger
		userMemberGuild := GetUserMemberGuild(session, previousMemberGuildId)
		memberMasterId := userMemberGuild.MemberMasterId
		biasRanking(&resp.MemberGuildTopStatus.MemberGuildRankingResultAnimationInfo, memberMasterId)
		resp.MemberGuildTopStatus.MemberGuildRankingAnimationInfo =
			resp.MemberGuildTopStatus.MemberGuildRankingResultAnimationInfo

		var userRank int64
		var err error
		if userMemberGuild.TotalPoint > 0 {
			userRank, err = session.Db.Table("u_member_guild").Where("member_master_id = ? AND member_guild_id = ? AND total_point > ?",
				memberMasterId, previousMemberGuildId, userMemberGuild.TotalPoint).OrderBy("total_point DESC").
				Limit(int(session.Gamedata.MemberGuildRankingRewardInside[memberMasterId].RankNumberLimit)).Count()
			utils.CheckErr(err)
			userRank++
			if int32(userRank) > session.Gamedata.MemberGuildRankingRewardInside[memberMasterId].RankNumberLimit {
				// having rally power and rank 0 mean unranked
				userRank = 0
			}
		}

		resp.MemberGuildTopStatus.MemberGuildInfo = client.MemberGuildTopInfo{
			MemberGuildUserRankingOrder: int32(userRank),
			MemberGuildUserRankingPoint: userMemberGuild.TotalPoint,
			// the other fields are unnecessary here
		}
		if hasResultTrigger {
			insideContent := session.Gamedata.MemberGuildRankingRewardInside[userMemberGuild.MemberMasterId].
				GetRewardContent(int32(userRank))
			if userRank > 0 {
				user_present.AddPresentWithDuration(session, client.PresentItem{
					Content:          *insideContent,
					PresentRouteType: enum.PresentRouteTypeMemberGuildInsideRankingReward,
					PresentRouteId:   generic.NewNullable(memberMasterId),
					ParamClient:      generic.NewNullable(fmt.Sprint(userRank)),
				}, user_present.Duration30Days)
			} else {
				user_present.AddPresentWithDuration(session, client.PresentItem{
					Content:          *insideContent,
					PresentRouteType: enum.PresentRouteTypeMemberGuildInsideRankingReward,
					PresentRouteId:   generic.NewNullable(memberMasterId),
				}, user_present.Duration30Days)
			}

			guildRank := int32(1)
			for _, info := range resp.MemberGuildTopStatus.MemberGuildRankingResultAnimationInfo.Slice {
				if info.MemberMasterId == memberMasterId {
					guildRank = info.MemberGuildRankingOrder
					break
				}
			}
			outsideContent := session.Gamedata.MemberGuildRankingRewardOutside[userMemberGuild.MemberMasterId].
				GetRewardContent(guildRank)

			user_present.AddPresentWithDuration(session, client.PresentItem{
				Content:          *outsideContent,
				PresentRouteType: enum.PresentRouteTypeMemberGuildOutsideRankingReward,
				PresentRouteId:   generic.NewNullable(memberMasterId),
				ParamClient:      generic.NewNullable(fmt.Sprint(guildRank)),
			}, user_present.Duration30Days)
		}
	} else {
		resp.MemberGuildTopStatus.MemberGuildRankingAnimationInfo =
			FetchMemberGuildRankingAnimationInfos(session, currentMemberGuildId)
		if session.UserStatus.MemberGuildMemberMasterId.HasValue {
			memberMasterId := session.UserStatus.MemberGuildMemberMasterId.Value
			biasRanking(&resp.MemberGuildTopStatus.MemberGuildRankingAnimationInfo, memberMasterId)
			currentUserTotalPoint := GetCurrentUserMemberGuildTotalPoint(session)
			var rank int64
			var err error
			if currentUserTotalPoint > 0 {
				rank, err = session.Db.Table("u_member_guild").Where("member_master_id = ? AND member_guild_id = ? AND total_point > ?",
					memberMasterId, currentMemberGuildId, currentUserTotalPoint).OrderBy("total_point DESC").
					Limit(int(session.Gamedata.MemberGuildRankingRewardInside[memberMasterId].RankNumberLimit)).Count()
				utils.CheckErr(err)
				rank++
				if int32(rank) > session.Gamedata.MemberGuildRankingRewardInside[memberMasterId].RankNumberLimit {
					rank = 0
				}
			}

			resp.MemberGuildTopStatus.MemberGuildInfo = client.MemberGuildTopInfo{
				MemberGuildUserRankingOrder: int32(rank),
				MemberGuildUserRankingPoint: currentUserTotalPoint,
				DailyCoopPoint:              GetDailyCoopPoint(session),
				CoopRewardPeriodAt:          generic.NewNullable(utils.NextMidDay(session.Time).Unix()),
			}
		}
	}
	return resp
}
