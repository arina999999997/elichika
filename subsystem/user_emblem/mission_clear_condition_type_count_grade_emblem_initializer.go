package user_emblem

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
	"elichika/utils"
)

func missionClearConditionTypeCountGradeEmblemInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]
	if mission.MissionClearConditionParam1 == nil { // no grade requirement
		count, err := session.Db.Table("u_emblem").Where("user_id = ?", session.UserId).Count()
		utils.CheckErr(err)
		userMission.MissionCount = int32(count)
	} else { // manually check, this might be a bit slow but whatever
		owned := []int32{}
		err := session.Db.Table("u_emblem").Where("user_id = ?", session.UserId).Cols("emblem_m_id").Find(&owned)
		utils.CheckErr(err)
		userMission.MissionCount = 0
		for _, emblemId := range owned {
			if session.Gamedata.Emblem[emblemId].Grade >= *mission.MissionClearConditionParam1 {
				userMission.MissionCount++
			}
		}
	}
	userMission.IsCleared = userMission.MissionCount >= mission.MissionClearConditionCount
	return userMission
}

func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeCountGradeEmblem, missionClearConditionTypeCountGradeEmblemInitializer)
}
