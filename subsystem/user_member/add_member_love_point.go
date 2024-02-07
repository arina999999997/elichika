package user_member

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_info_trigger"
	"elichika/subsystem/user_present"
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
	if oldLoveLevel < member.LoveLevel {
		masterMember := session.Gamedata.Member[memberId]
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
	UpdateMember(session, member)
	return point
}
