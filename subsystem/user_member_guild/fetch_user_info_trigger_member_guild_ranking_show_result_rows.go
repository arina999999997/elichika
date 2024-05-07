package user_member_guild

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
)

func FetchUserInfoTriggerMemberGuildRankingShowResultRows(session *userdata.Session,
	resultList *generic.List[client.UserInfoTriggerMemberGuildRankingShowResultRow]) {
	triggerId, hasReward := CheckPreviousMemberGuildReward(session)
	if !hasReward {
		return
	}

	previousMemberGuildId := GetCurrentMemberGuildId(session) - 1
	_, resultAt := GetMemberGuildStartAndEnd(session, previousMemberGuildId)
	limitAt := resultAt + session.Gamedata.MemberGuildPeriod.OneCycleSecs
	resultList.Append(client.UserInfoTriggerMemberGuildRankingShowResultRow{
		TriggerId:     triggerId,
		MemberGuildId: previousMemberGuildId,
		ResultAt:      resultAt,
		EndAt:         limitAt,
	})
}
