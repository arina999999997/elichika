package user_member

import (
	"elichika/config"
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"
)

func TapLovePoint(session *userdata.Session, memberMasterId int32) {
	amount := *config.Conf.TapBondGain
	if session.UserStatus.TutorialPhase == enum.TutorialPhaseLovePointUp {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseTrainingLevelUp
		// too much love point can freeze the tutorial, so we give out the default amount
		amount = 20
	}
	AddMemberLovePoint(session, memberMasterId, amount)
	if config.Conf.ResourceConfig().ConsumeNaviTap {
		if session.UserStatus.NaviTapRecoverAt <= session.Time.Unix() {
			session.UserStatus.NaviTapCount = 0
			session.UserStatus.NaviTapRecoverAt = utils.BeginOfNextDay(session.Time).Unix()
		}
		session.UserStatus.NaviTapCount++
	}
}
