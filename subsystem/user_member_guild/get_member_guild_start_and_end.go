package user_member_guild

import (
	"elichika/userdata"
)

// the year is based on the end of the period, not beginning
func GetMemberGuildStartAndEnd(session *userdata.Session, memberGuildId int32) (int64, int64) {
	endTime := session.Gamedata.MemberGuildPeriod.StartAt + session.Gamedata.MemberGuildPeriod.OneCycleSecs*int64(memberGuildId)
	startTime := endTime - session.Gamedata.MemberGuildPeriod.OneCycleSecs
	return startTime, endTime
}
