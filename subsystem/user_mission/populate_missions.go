package user_mission

import (
	"elichika/enum"
	"elichika/userdata"
)

func populateMissions(session *userdata.Session) {
	for _, mission := range session.Gamedata.MissionByTerm[enum.MissionTermFree] {
		if (mission.StartAt > session.Time.Unix()) || (mission.EndAt < session.Time.Unix()) {
			continue
		}

		userMission := getUserMission(session, mission.Id)
		if userMission.MissionMId != 0 { // valid mission that we can have
			session.UserModel.UserMissionByMissionId.Set(mission.Id, userMission)
		}
	}
}
