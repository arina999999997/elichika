package user_mission

import (
	"elichika/client"
	"elichika/gamedata"
	"elichika/userdata"
)

func GetMasterMission(session *userdata.Session, mission any) *gamedata.Mission {
	switch mission := mission.(type) {
	case client.UserMission:
		return session.Gamedata.Mission[mission.MissionMId]
	case client.UserDailyMission:
		return session.Gamedata.Mission[mission.MissionMId]
	case client.UserWeeklyMission:
		return session.Gamedata.Mission[mission.MissionMId]
	default:
		panic("not supported")
	}
}
