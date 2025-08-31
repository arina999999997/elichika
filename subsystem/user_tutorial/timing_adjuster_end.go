package user_tutorial

import (
	"elichika/enum"
	"elichika/log"
	"elichika/userdata"
)

func TimingAdjusterEnd(session *userdata.Session) {
	if session.UserStatus.TutorialPhase != enum.TutorialPhaseTimingAdjuster {
		log.Panic("Unexpected tutorial phase")
	}
	session.UserStatus.TutorialPhase = enum.TutorialPhaseFavoriateMember
}
