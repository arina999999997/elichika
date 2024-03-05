package user_mission

import (
	"elichika/enum"
	"elichika/userdata"
)

func populateWeeklyMissions(session *userdata.Session) {
	for _, mission := range session.Gamedata.MissionByTerm[enum.MissionTermWeekly] {
		if (mission.StartAt > session.Time.Unix()) || (mission.EndAt < session.Time.Unix()) {
			continue
		}

		userWeeklyMission := getUserWeeklyMission(session, mission.Id)
		if userWeeklyMission.MissionMId != 0 { // valid mission that we can have
			session.UserModel.UserWeeklyMissionByMissionId.Set(mission.Id, userWeeklyMission)
		}
	}
}
