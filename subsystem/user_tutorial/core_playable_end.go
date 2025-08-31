package user_tutorial

import (
	"elichika/enum"
	"elichika/log"
	"elichika/userdata"
)

func CorePlayableEnd(session *userdata.Session) {
	if session.UserStatus.TutorialPhase != enum.TutorialPhaseCorePlayable {
		log.Panic("Unexpected tutorial phase")
	}
	session.UserStatus.TutorialPhase = enum.TutorialPhaseTimingAdjuster
}
