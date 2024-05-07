package user_member_guild

import (
	"elichika/userdata"
)

func GetCurrentUserMemberGuildTotalPoint(session *userdata.Session) int32 {
	userMemberGuild := GetCurrentUserMemberGuild(session)
	return userMemberGuild.SupportPoint + userMemberGuild.LovePoint + userMemberGuild.VoltagePoint
}
