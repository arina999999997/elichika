package user_member_guild

import (
	"elichika/userdata"
)

// note that this is the Id the client prefer, using the wrong id will lead to the client ignoring stuff.
// for cross server playing, it's better to just update the start time of them to sync up
// alternatively, just run the server for a single masterdata
func GetCurrentMemberGuildId(session *userdata.Session) int32 {
	return int32((session.Time.Unix()-session.Gamedata.MemberGuildPeriod.StartAt)/
		int64(session.Gamedata.MemberGuildPeriod.OneCycleSecs)) + 1
}
