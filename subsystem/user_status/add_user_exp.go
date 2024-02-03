package user_status

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"

	"elichika/subsystem/user_info_trigger"
)

func AddUserExp(session *userdata.Session, exp int32) {
	session.UserStatus.Exp += exp
	trigger := client.UserInfoTriggerBasic{
		InfoTriggerType: enum.InfoTriggerTypeUserLevelUp,
		ParamInt:        generic.NewNullable(session.UserStatus.Rank),
	}
	isRankedUp := false
	for session.UserStatus.Rank < session.Gamedata.UserRankMax {
		nextRank := session.Gamedata.UserRank[session.UserStatus.Rank+1]
		if session.UserStatus.Exp >= nextRank.Exp {
			isRankedUp = true
			session.UserStatus.Rank++
			AddUserLp(session, nextRank.MaxLp)
			AddUserAccessoryLimit(session, nextRank.AdditionalAccessoryLimit)
		} else {
			break
		}
	}
	if isRankedUp {
		user_info_trigger.AddTriggerBasic(session, trigger)
	}
}
