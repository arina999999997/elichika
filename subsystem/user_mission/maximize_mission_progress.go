package user_mission

import (
	"elichika/client"
	"elichika/userdata"
)

// maximize the mission progress and update it directly
// assuming we already filter through UpdateProgress first
func MaximizeMissionProgress(session *userdata.Session, mission any, count int32) {
	switch mission.(type) {
	case client.UserMission:
		userMission := mission.(client.UserMission)
		masterMission := session.Gamedata.Mission[userMission.MissionMId]
		if userMission.MissionCount < count {
			userMission.MissionCount = count
		}
		if (!userMission.IsCleared) && (userMission.MissionCount >= masterMission.MissionClearConditionCount) {
			userMission.IsCleared = true
		}
		session.UserModel.UserMissionByMissionId.Set(userMission.MissionMId, userMission)
	default:
		panic("not supported")
	}
}
