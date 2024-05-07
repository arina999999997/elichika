package user_member_guild

import (
	"elichika/userdata"
)

func AddLovePoint(session *userdata.Session, loveGained int32) int32 {
	if !IsMemberGuildRankingPeriod(session) {
		return 0
	}
	userMemberGuild := GetCurrentUserMemberGuild(session)
	pointGained := loveGained / session.Gamedata.MemberGuildConstant.LovePointCalculationNum
	pointAllowed := session.Gamedata.MemberGuildConstant.DailyLimitPoint - userMemberGuild.DailyLovePoint
	if pointGained > pointAllowed {
		pointGained = pointAllowed
	}
	if pointGained <= 0 {
		return 0
	}
	userMemberGuild.LovePoint += pointGained
	userMemberGuild.DailyLovePoint += pointGained
	UpdateUserMemberGuild(session, userMemberGuild)
	return pointGained
}
