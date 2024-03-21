package user_mission

import (
	"elichika/enum"
	"elichika/userdata"
)

func ClearMissionNewBadge(session *userdata.Session, missionTerm int32) {
	session.SendMissionDetail = true

	for _, mission := range session.Gamedata.MissionByTerm[missionTerm] {
		if (mission.StartAt > session.Time.Unix()) || (mission.EndAt < session.Time.Unix()) {
			continue
		}
		switch missionTerm {
		case enum.MissionTermDaily:
			userDailyMission := getUserDailyMission(session, mission.Id)
			if (userDailyMission.MissionMId != 0) && userDailyMission.IsNew {
				userDailyMission.IsNew = false
				session.UserModel.UserDailyMissionByMissionId.Set(mission.Id, userDailyMission)
			}
		case enum.MissionTermWeekly:
			userWeeklyMission := getUserWeeklyMission(session, mission.Id)
			if (userWeeklyMission.MissionMId != 0) && userWeeklyMission.IsNew {
				userWeeklyMission.IsNew = false
				userWeeklyMission.NewExpiredAt = session.Time.Unix()
				session.UserModel.UserWeeklyMissionByMissionId.Set(mission.Id, userWeeklyMission)
			}
		default:
			userMission := getUserMission(session, mission.Id)
			if (userMission.MissionMId != 0) && userMission.IsNew {
				userMission.IsNew = false
				userMission.NewExpiredAt = session.Time.Unix()
				session.UserModel.UserMissionByMissionId.Set(mission.Id, userMission)
			}
		}
	}
}
