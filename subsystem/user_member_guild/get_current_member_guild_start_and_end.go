package user_member_guild

import (
	"elichika/userdata"
)

// the year is based on the end of the period, not beginning
func GetCurrentMemberGuildStartAndEnd(session *userdata.Session) (int64, int64) {
	startTime, endTime := GetMemberGuildStartAndEnd(session, GetCurrentMemberGuildId(session))
	return startTime, endTime
}

func GetCurrentMemberGuildStart(session *userdata.Session) int64 {
	startTime, _ := GetCurrentMemberGuildStartAndEnd(session)
	return startTime
}

func GetCurrentMemberGuildEnd(session *userdata.Session) int64 {
	_, endTime := GetCurrentMemberGuildStartAndEnd(session)
	return endTime
}
