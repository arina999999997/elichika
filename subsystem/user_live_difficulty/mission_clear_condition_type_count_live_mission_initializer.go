package user_live_difficulty

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
	"elichika/utils"
)

func missionClearConditionTypeCountLiveMissionInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]

	userMission.MissionCount = 0
	count, err := session.Db.Table("u_live_difficulty").
		Where("user_id = ? AND cleared_difficulty_achievement_1 = 1", session.UserId).Count()
	utils.CheckErr(err)
	userMission.MissionCount += int32(count)
	count, err = session.Db.Table("u_live_difficulty").
		Where("user_id = ? AND cleared_difficulty_achievement_2 = 2", session.UserId).Count()
	utils.CheckErr(err)
	userMission.MissionCount += int32(count)
	count, err = session.Db.Table("u_live_difficulty").
		Where("user_id = ? AND cleared_difficulty_achievement_3 = 3", session.UserId).Count()
	utils.CheckErr(err)
	userMission.MissionCount += int32(count)

	userMission.IsCleared = userMission.MissionCount >= mission.MissionClearConditionCount
	return userMission
}

func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeCountLiveMission, missionClearConditionTypeCountLiveMissionInitializer)
}
