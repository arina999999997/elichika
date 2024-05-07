package user_member_guild

import (
	"elichika/userdata"
)

func IsMemberGuildRankingPeriod(session *userdata.Session) bool {
	period := session.Gamedata.MemberGuildPeriod
	currentPeriodTime := session.Time.Unix() - GetCurrentMemberGuildStart(session)
	return (currentPeriodTime >= period.RankingStartSecs) && (currentPeriodTime < period.RankingEndSecs)
}

func IsBeforeMemberGuildRankingPeriod(session *userdata.Session) bool {
	period := session.Gamedata.MemberGuildPeriod
	currentPeriodTime := session.Time.Unix() - GetCurrentMemberGuildStart(session)
	return currentPeriodTime < period.RankingStartSecs
}
