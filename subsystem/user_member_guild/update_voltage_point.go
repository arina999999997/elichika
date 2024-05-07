package user_member_guild

import (
	"elichika/userdata"
)

func UpdateVoltagePoint(session *userdata.Session, liveId, voltagePoint int32) (int32, bool) {
	if liveId != session.Gamedata.MemberGuildChallengeLive.GetLiveId(GetCurrentMemberGuildId(session)) {
		// not the current ranking song
		return 0, false
	}
	userMemberGuild := GetCurrentUserMemberGuild(session)
	if userMemberGuild.MaxVoltage < int64(voltagePoint) {
		userMemberGuild.MaxVoltage = int64(voltagePoint)
		userMemberGuild.VoltagePoint = voltagePoint / session.Gamedata.MemberGuildConstant.VoltageCalculationNum
		UpdateUserMemberGuild(session, userMemberGuild)
		return userMemberGuild.VoltagePoint, true
	} else {
		return 0, false
	}
}
