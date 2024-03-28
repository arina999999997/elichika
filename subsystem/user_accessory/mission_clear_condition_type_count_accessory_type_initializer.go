package user_accessory

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
	"elichika/utils"
)

func missionClearConditionTypeCountAccessoryTypeInitializer(session *userdata.Session, userMission client.UserMission) client.UserMission {
	mission := session.Gamedata.Mission[userMission.MissionMId]
	accessories := []client.UserAccessory{}

	err := session.Db.Table("u_accessory").Where("user_id = ?", session.UserId).Find(&accessories)
	utils.CheckErr(err)

	has := map[int32]bool{}
	for _, accessory := range accessories {
		has[accessory.AccessoryMasterId] = true
	}
	userMission.MissionCount = int32(len(has))
	userMission.IsCleared = userMission.MissionCount >= mission.MissionClearConditionCount
	return userMission
}

func init() {
	user_mission.AddMissionInitializer(enum.MissionClearConditionTypeCountAccessoryType, missionClearConditionTypeCountAccessoryTypeInitializer)
}
