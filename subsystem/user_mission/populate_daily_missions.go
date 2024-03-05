package user_mission

import (
	"elichika/enum"
	"elichika/userdata"
)

func populateDailyMissions(session *userdata.Session) {
	for _, mission := range session.Gamedata.MissionByTerm[enum.MissionTermDaily] {
		if (mission.StartAt > session.Time.Unix()) || (mission.EndAt < session.Time.Unix()) {
			continue
		}

		userDailyMission := getUserDailyMission(session, mission.Id)
		if userDailyMission.MissionMId != 0 { // valid mission that we can have
			session.UserModel.UserDailyMissionByMissionId.Set(mission.Id, userDailyMission)
		}
	}
}
