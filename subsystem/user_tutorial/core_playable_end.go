package user_tutorial

import (
	"elichika/enum"
	"elichika/userdata"
)

func CorePlayableEnd(session *userdata.Session) {
	if session.UserStatus.TutorialPhase != enum.TutorialPhaseCorePlayable {
		panic("Unexpected tutorial phase")
	}
	session.UserStatus.TutorialPhase = enum.TutorialPhaseTimingAdjuster
}
