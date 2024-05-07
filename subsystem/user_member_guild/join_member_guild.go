package user_member_guild

import (
	"elichika/generic"
	"elichika/userdata"
)

func JoinMemberGuild(session *userdata.Session, memberMasterId int32) {
	session.UserStatus.MemberGuildMemberMasterId = generic.NewNullable(memberMasterId)
	session.UserStatus.MemberGuildLastUpdatedAt = session.Time.Unix()
	// reset the current member guild if it's already created
	userMemberGuild := GetCurrentUserMemberGuild(session)
	userMemberGuild.MemberMasterId = memberMasterId
	userMemberGuild.DailyLovePoint = 0
	userMemberGuild.DailySupportPoint = 0
	userMemberGuild.SupportPoint = 0
	userMemberGuild.LovePoint = 0
	userMemberGuild.VoltagePoint = 0
	userMemberGuild.MaxVoltage = 0
	UpdateUserMemberGuild(session, userMemberGuild)
}
