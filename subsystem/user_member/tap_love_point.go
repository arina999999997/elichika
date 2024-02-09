package user_member

import (
	"elichika/config"
	"elichika/enum"
	"elichika/userdata"
)

func TapLovePoint(session *userdata.Session, memberMasterId int32) {
	if session.UserStatus.TutorialPhase == enum.TutorialPhaseLovePointUp {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseTrainingLevelUp
		// too much love point can freeze the tutorial, so we give out the default amount
		AddMemberLovePoint(session, memberMasterId, 20)
	} else {
		AddMemberLovePoint(session, memberMasterId, *config.Conf.TapBondGain)
	}
}
