package user_mission

import (
	"elichika/client"
	"elichika/enum"
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
		if (!userMission.IsCleared) && (userMission.MissionCount >= masterMission.MissionClearConditionCount) {
			userMission.IsCleared = true
		}
		session.UserModel.UserMissionByMissionId.Set(userMission.MissionMId, userMission)
	case client.UserDailyMission:
		userDailyMission := mission.(client.UserDailyMission)
		masterMission := session.Gamedata.Mission[userDailyMission.MissionMId]
		userDailyMission.MissionCount += count
		if (!userDailyMission.IsCleared) &&
			((userDailyMission.MissionCount - userDailyMission.MissionStartCount) >= masterMission.MissionClearConditionCount) {
			userDailyMission.IsCleared = true
			session.UserModel.UserDailyMissionByMissionId.Set(userDailyMission.MissionMId, userDailyMission)
			UpdateProgress(session, enum.MissionClearConditionTypeCompleteDaily, nil, nil,
				func(session *userdata.Session, missionList []any, _ ...any) {
					for _, otherMission := range missionList {
						otherMasterMission := session.Gamedata.Mission[otherMission.(client.UserDailyMission).MissionMId]
						if (masterMission.PickupType == nil) != (otherMasterMission.PickupType == nil) {
							continue
						}
						if (masterMission.PickupType == nil) || (*masterMission.PickupType == *otherMasterMission.PickupType) {
							AddMissionProgress(session, otherMission, int32(1))
						}
					}
				})
		} else {
			session.UserModel.UserDailyMissionByMissionId.Set(userDailyMission.MissionMId, userDailyMission)
		}
	case client.UserWeeklyMission:
		userWeeklyMission := mission.(client.UserWeeklyMission)
		masterMission := session.Gamedata.Mission[userWeeklyMission.MissionMId]
		userWeeklyMission.MissionCount += count
		if (!userWeeklyMission.IsCleared) &&
			((userWeeklyMission.MissionCount - userWeeklyMission.MissionStartCount) >= masterMission.MissionClearConditionCount) {
			userWeeklyMission.IsCleared = true
			session.UserModel.UserWeeklyMissionByMissionId.Set(userWeeklyMission.MissionMId, userWeeklyMission)
			UpdateProgress(session, enum.MissionClearConditionTypeCompleteWeekly, nil, nil, AddProgressHandler, int32(1))
		} else {
			session.UserModel.UserWeeklyMissionByMissionId.Set(userWeeklyMission.MissionMId, userWeeklyMission)
		}
	default:
		panic("not supported")
	}
}
