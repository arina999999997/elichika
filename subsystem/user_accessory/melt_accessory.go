package user_accessory

import (
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
)

func MeltAccessory(session *userdata.Session, userAccessoryId int64) {
	accessory := GetUserAccessory(session, userAccessoryId)
	user_content.AddContent(session, session.Gamedata.Accessory[accessory.AccessoryMasterId].MeltGroup[accessory.Grade].Reward)
	DeleteUserAccessory(session, userAccessoryId)
	// mission
	user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountAccessoryMelt, nil, nil,
		user_mission.AddProgressHandler, int32(1))
}
