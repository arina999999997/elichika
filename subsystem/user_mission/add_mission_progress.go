package user_mission

import (
	"elichika/client"
	"elichika/userdata"
)

// add the mission progress and update it directly
// assuming we already filter through UpdateProgress first
func AddMissionProgress(session *userdata.Session, mission any, count int32) {
	switch mission.(type) {
	case client.UserMission:
		userMission := mission.(client.UserMission)
		masterMission := session.Gamedata.Mission[userMission.MissionMId]
		userMission.MissionCount += count
		userMission.IsCleared = userMission.MissionCount >= masterMission.MissionClearConditionCount
		session.UserModel.UserMissionByMissionId.Set(userMission.MissionMId, userMission)
	case client.UserDailyMission:
		userDailyMission := mission.(client.UserDailyMission)
		masterMission := session.Gamedata.Mission[userDailyMission.MissionMId]
		userDailyMission.MissionCount += count
		userDailyMission.IsCleared = (userDailyMission.MissionCount - userDailyMission.MissionStartCount) >= masterMission.MissionClearConditionCount
		session.UserModel.UserDailyMissionByMissionId.Set(userDailyMission.MissionMId, userDailyMission)
	case client.UserWeeklyMission:
		userWeeklyMission := mission.(client.UserWeeklyMission)
		masterMission := session.Gamedata.Mission[userWeeklyMission.MissionMId]
		userWeeklyMission.MissionCount += count
		userWeeklyMission.IsCleared = (userWeeklyMission.MissionCount - userWeeklyMission.MissionStartCount) >= masterMission.MissionClearConditionCount
		session.UserModel.UserWeeklyMissionByMissionId.Set(userWeeklyMission.MissionMId, userWeeklyMission)

	default:
		panic("not supported")
	}
}
