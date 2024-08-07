package user_accessory

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
	"elichika/utils"
)

// return true if some has to be in present box
func accessoryContentHandler(session *userdata.Session, content *client.Content) any {
	// technically this counting method is incorrect but it sohuld be fine in this context (present / adding accessory)
	currentAccessoryAmount, err := session.Db.Table("u_accessory").Where("user_id = ?", session.UserId).
		Count(&client.UserAccessory{})
	utils.CheckErr(err)
	currentAccessoryAmount += int64(session.UserModel.UserAccessoryByUserAccessoryId.Size())

	masterAccessory := session.Gamedata.Accessory[content.ContentId]
	for content.ContentAmount > 0 {
		currentAccessoryAmount++
		if int32(currentAccessoryAmount) > session.UserStatus.AccessoryBoxAdditional+AccessoryBoxDefaultLimit {
			// can't take anymore
			return true
		}
		content.ContentAmount--
		accessory := GetUserAccessory(session, session.NextUniqueId())
		accessory.AccessoryMasterId = masterAccessory.Id
		accessory.Level = 1
		accessory.Exp = 0
		accessory.Grade = 0
		accessory.Attribute = masterAccessory.Attribute
		if masterAccessory.Grade[0].PassiveSkill1MasterId != nil {
			accessory.PassiveSkill1Id = generic.NewNullable(*masterAccessory.Grade[0].PassiveSkill1MasterId)
		}
		if masterAccessory.Grade[0].PassiveSkill2MasterId != nil {
			accessory.PassiveSkill2Id = generic.NewNullable(*masterAccessory.Grade[0].PassiveSkill2MasterId)
		}
		UpdateUserAccessory(session, accessory)

		// mission
		user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountAccessory, nil, nil,
			user_mission.AddProgressHandler, int32(1))
	}
	return false
}

func init() {
	user_content.AddContentHandler(enum.ContentTypeAccessory, accessoryContentHandler)
}
