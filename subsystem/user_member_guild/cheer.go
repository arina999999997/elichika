package user_member_guild

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/item"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_present"
	"elichika/userdata"
)

func Cheer(session *userdata.Session, count generic.Nullable[int32]) response.CheerMemberGuildResponse {
	resp := response.CheerMemberGuildResponse{
		UserModelDiff: &session.UserModel,
	}

	// this is the same with FetchMemberGuildTop
	rank := int32(0)
	for _, member := range session.Gamedata.Member {
		rank++
		resp.MemberGuildTopStatus.MemberGuildRankingAnimationInfo.Append(
			client.MemberGuildRankingAnimationInfo{
				MemberMasterId:          member.Id,
				MemberGuildRankingOrder: rank,
				MemberGuildRankingPoint: 100000 - rank*1000,
			})
	}

	// TODO(member_guild): Award cheer points when we implement it
	if !count.HasValue {
		count.Value = 1
		// TODO(member_guild): Mark this as used?
	} else {
		user_content.RemoveContent(session, item.RallyMegaphone.Amount(-count.Value))
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
	return resp
}
