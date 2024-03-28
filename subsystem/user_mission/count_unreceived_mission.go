package user_mission

import (
	"elichika/userdata"
)

// this is used in fetch bootstrap, we will fetch the mission again then count
func CountUnreceivedMission(session *userdata.Session) int32 {
	populateMissions(session)
	populateDailyMissions(session)
	populateWeeklyMissions(session)
	count := int32(0)
	for _, mission := range session.UserModel.UserMissionByMissionId.Map {
		if mission.IsCleared && (!mission.IsReceivedReward) {
			count++
		}
	}
	for _, mission := range session.UserModel.UserDailyMissionByMissionId.Map {
		if mission.IsCleared && (!mission.IsReceivedReward) {
			count++
		}
	}
	for _, mission := range session.UserModel.UserWeeklyMissionByMissionId.Map {
		if mission.IsCleared && (!mission.IsReceivedReward) {
			count++
		}
	}
	return count
}
