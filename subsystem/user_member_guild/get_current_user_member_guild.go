package user_member_guild

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_info_trigger"
	"elichika/subsystem/user_present"
	"elichika/userdata"
	"elichika/utils"
)

func GetCurrentUserMemberGuild(session *userdata.Session) client.UserMemberGuild {
	memberGuildId := GetCurrentMemberGuildId(session)
	_, exist := session.UserModel.UserMemberGuildById.Map[memberGuildId]
	if exist {
		return *session.UserModel.UserMemberGuildById.Map[memberGuildId]
	}
	userMemberGuild := client.UserMemberGuild{}
	exist, err := session.Db.Table("u_member_guild").Where("user_id = ? AND member_guild_id = ?",
		session.UserId, GetCurrentMemberGuildId(session)).Get(&userMemberGuild)
	utils.CheckErr(err)
	if !exist {
		userMemberGuild = client.UserMemberGuild{
			MemberGuildId:            memberGuildId,
			MemberMasterId:           session.UserModel.UserStatus.MemberGuildMemberMasterId.Value,
			SupportPointCountResetAt: utils.BeginOfCurrentHalfDay(session.Time).Unix(),
		}
		UpdateUserMemberGuild(session, userMemberGuild)
	}
	if int64(userMemberGuild.DailySupportPointResetAt) <= session.Time.Unix() {
		// reset the daily point
		// TODO(extra): This system can't handle more than 1 rally goal, not sure if the client can handle it
		previousDayPoint := GetPreviousDailyCoopPoint(session, utils.CurrentMidDay(session.Time).Unix())
		memberMasterId := session.UserModel.UserStatus.MemberGuildMemberMasterId.Value
		reward := session.Gamedata.MemberGuildPointClearReward[memberMasterId]
		if (previousDayPoint >= reward.TargetPoint) &&
			(userMemberGuild.DailySupportPoint+userMemberGuild.DailyLovePoint > 0) &&
			(utils.CurrentMidDay(session.Time).Unix() == int64(userMemberGuild.DailySupportPointResetAt)) {
			user_info_trigger.AddTriggerBasic(session,
				client.UserInfoTriggerBasic{
					InfoTriggerType: enum.InfoTriggerTypeMemberGuildTopRewardReceived,
					ParamInt:        generic.NewNullable(reward.TargetPoint),
				})
			user_present.AddPresentWithDuration(session,
				client.PresentItem{
					Content:          reward.Content,
					PresentRouteType: enum.PresentRouteTypeMemberGuildPointClearReward,
					PresentRouteId:   generic.NewNullable(memberMasterId),
				}, user_present.Duration30Days)
		}
		userMemberGuild.DailySupportPoint = 0
		userMemberGuild.DailyLovePoint = 0
		userMemberGuild.DailySupportPointResetAt = int32(utils.NextMidDay(session.Time).Unix())
		userMemberGuild.DailyLovePointResetAt = userMemberGuild.DailySupportPointResetAt
		UpdateUserMemberGuild(session, userMemberGuild)
	}
	return userMemberGuild
}
