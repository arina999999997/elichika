package user_mission

import (
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
)

// fetch the mission ids without populating the user model field
func FetchMissionIds(session *userdata.Session) generic.List[int32] {
	missionIds := generic.List[int32]{}

	for _, mission := range session.Gamedata.MissionByTerm[enum.MissionTermDaily] {
		if (mission.StartAt > session.Time.Unix()) || (mission.EndAt < session.Time.Unix()) {
			continue
		}
		userDailyMission := getUserDailyMission(session, mission.Id)
		if userDailyMission.MissionMId != 0 { // valid mission that we can have
			missionIds.Append(userDailyMission.MissionMId)
			if userDailyMission.IsNew {
				session.UserModel.UserDailyMissionByMissionId.Set(mission.Id, userDailyMission)
			}
		}
	}

	for _, mission := range session.Gamedata.MissionByTerm[enum.MissionTermWeekly] {
		if (mission.StartAt > session.Time.Unix()) || (mission.EndAt < session.Time.Unix()) {
			continue
		}
		userWeeklyMission := getUserWeeklyMission(session, mission.Id)
		if userWeeklyMission.MissionMId != 0 { // valid mission that we can have
			missionIds.Append(userWeeklyMission.MissionMId)
		}
		if userWeeklyMission.IsNew {
			session.UserModel.UserWeeklyMissionByMissionId.Set(mission.Id, userWeeklyMission)
		}
	}

	for _, mission := range session.Gamedata.MissionByTerm[enum.MissionTermFree] {
		if (mission.StartAt > session.Time.Unix()) || (mission.EndAt < session.Time.Unix()) {
			continue
		}
		userMission := getUserMission(session, mission.Id)
		if userMission.MissionMId != 0 { // valid mission that we can have
			missionIds.Append(userMission.MissionMId)
		}
		if userMission.IsNew {
			session.UserModel.UserMissionByMissionId.Set(mission.Id, userMission)
		}
	}

	return missionIds
}
