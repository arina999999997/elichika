package user_member

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_info_trigger"
	"elichika/subsystem/user_present"
	"elichika/subsystem/user_unlock_scene"
	"elichika/userdata"

	"fmt"
)

// add love point and return the love point added (in case maxed out)
func AddMemberLovePoint(session *userdata.Session, memberId, point int32) int32 {
	member := GetMember(session, memberId)
	if point > member.LovePointLimit-member.LovePoint {
		point = member.LovePointLimit - member.LovePoint
	}
	member.LovePoint += point

	oldLoveLevel := member.LoveLevel
	member.LoveLevel = session.Gamedata.LoveLevelFromLovePoint(member.LovePoint)
	// unlock bond stories, unlock bond board

	masterMember := session.Gamedata.Member[memberId]
	if oldLoveLevel < member.LoveLevel {
		for loveLevel := oldLoveLevel + 1; loveLevel <= member.LoveLevel; loveLevel++ {
			for _, reward := range masterMember.LoveLevelRewards[loveLevel] {
				user_present.AddPresent(session, client.PresentItem{
					Content:          reward,
					PresentRouteType: enum.PresentRouteTypeLoveLevelUp,
					PresentRouteId:   generic.NewNullable(masterMember.LoveLevelRewardIds[loveLevel]),
					ParamClient:      generic.NewNullable(fmt.Sprint(member.MemberMasterId)),
				})
			}
		}
		user_info_trigger.AddTriggerMemberLoveLevelUp(session, client.UserInfoTriggerMemberLoveLevelUp{
			MemberMasterId:  memberId,
			BeforeLoveLevel: member.LoveLevel - 1})

		UnlockNewLovePanel(session, memberId, oldLoveLevel, member.LoveLevel)
	}

	// also award previous reward if we missed any
	// TODO(final): this is only necessary for updating users, and should be removed once the server is "finalized"
	for loveLevel := int32(1); loveLevel <= oldLoveLevel; loveLevel++ {
		for _, reward := range masterMember.LoveLevelRewards[loveLevel] {
			if session.Gamedata.ContentType[reward.ContentType].IsUnique {
				user_present.AddPresent(session, client.PresentItem{
					Content:          reward,
					PresentRouteType: enum.PresentRouteTypeLoveLevelUp,
					PresentRouteId:   generic.NewNullable(masterMember.LoveLevelRewardIds[loveLevel]),
					ParamClient:      generic.NewNullable(fmt.Sprint(member.MemberMasterId)),
				})
			}
		}
	}

	// special behavior to make sure no one get locked out
	if member.LoveLevel >= 10 {
		if user_unlock_scene.HasUnlockScene(session, enum.UnlockSceneTypeMemberGuild) {
			user_unlock_scene.UnlockScene(session, enum.UnlockSceneTypeMemberGuild, enum.UnlockSceneStatusOpened)
		} else {
			user_unlock_scene.UnlockScene(session, enum.UnlockSceneTypeMemberGuild, enum.UnlockSceneStatusOpen)
		}
	}
	UpdateMember(session, member)
	return point
}
