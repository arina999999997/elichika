package user_accessory

import (
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func MeltAccessory(session *userdata.Session, userAccessoryId int64) {
	accessory := GetUserAccessory(session, userAccessoryId)
	user_content.AddContent(session, session.Gamedata.Accessory[accessory.AccessoryMasterId].MeltGroup[accessory.Grade].Reward)
	DeleteUserAccessory(session, userAccessoryId)
}
