package user_mission

import (
	"elichika/enum"
	"elichika/userdata"
)

func finishedMission(session *userdata.Session, missionId int32) bool {
	mission := session.Gamedata.Mission[missionId]
	if mission == nil {
		return false
	}
	switch mission.Term {
	case enum.MissionTermDaily:
		return getUserDailyMission(session, missionId).IsReceivedReward
	case enum.MissionTermWeekly:
		return getUserWeeklyMission(session, missionId).IsReceivedReward
	default:
		return getUserMission(session, missionId).IsReceivedReward
	}
}
