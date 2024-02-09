package user_tutorial

import (
	"elichika/enum"
	"elichika/userdata"
)

func TimingAdjusterEnd(session *userdata.Session) {
	if session.UserStatus.TutorialPhase != enum.TutorialPhaseTimingAdjuster {
		panic("Unexpected tutorial phase")
	}
	session.UserStatus.TutorialPhase = enum.TutorialPhaseFavoriateMember
}
