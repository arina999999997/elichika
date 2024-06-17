package user_member_guild

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/config"
	"elichika/enum"
	"elichika/generic"
	"elichika/item"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_present"
	"elichika/userdata"
	"elichika/utils"
)

func CheerMemberGuild(session *userdata.Session, count generic.Nullable[int32]) response.CheerMemberGuildResponse {
	resp := response.CheerMemberGuildResponse{
		UserModelDiff: &session.UserModel,
	}
	userMemberGuild := GetCurrentUserMemberGuild(session)
	if !count.HasValue { // free chance
		userMemberGuild.SupportPointCountResetAt = utils.BeginOfNextHalfDay(session.Time).Unix()
		count.Value = 1
	} else {
		if config.Conf.ResourceConfig().ConsumeMemberCheerItem {
			user_content.RemoveContent(session, item.RallyMegaphone.Amount(count.Value))
		}
	}

	if IsMemberGuildRankingPeriod(session) {
		pointGain := session.Gamedata.MemberGuildConstant.SupportPoint * count.Value
		userMemberGuild.SupportPoint += pointGain
		userMemberGuild.DailySupportPoint += pointGain
	}

	for i := int32(0); i < count.Value; i++ {
		content := session.Gamedata.MemberGuildCheerReward[session.UserStatus.MemberGuildMemberMasterId.Value].GetRandomItem()
		user_present.AddPresent(session, client.PresentItem{
			Content:          content,
			PresentRouteType: enum.PresentRouteTypeMemberGuildSupportReward,
			PresentRouteId:   session.UserStatus.MemberGuildMemberMasterId, // this doesn't seems to be doing anything but the official server kept it
		})
		resp.Rewards.Append(content)
	}
	UpdateUserMemberGuild(session, userMemberGuild)
	return resp
}
