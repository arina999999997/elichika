package user_login

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
)

func missionClearConditionTypeCountLoginInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]
	userMission.MissionCount = session.UserStatus.LoginDays
	userMission.IsCleared = userMission.MissionCount >= mission.MissionClearConditionCount
	return userMission
}

func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeCountLogin, missionClearConditionTypeCountLoginInitializer)
}
