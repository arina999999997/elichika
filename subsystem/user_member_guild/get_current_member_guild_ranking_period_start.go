package user_member_guild

import (
	"elichika/userdata"
)

func GetCurrentMemberGuildRankingPeriodStart(session *userdata.Session) int64 {
	return session.Time.Unix() - (session.Time.Unix() % session.Gamedata.MemberGuildPeriod.OneCycleSecs)
}
