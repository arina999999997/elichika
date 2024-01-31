package user_accessory
import (
	"elichika/client"
	"elichika/client/response"
	"elichika/generic"
	"elichika/item"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func RarityUpAccessory(session *userdata.Session, userAccessoryId int64) response.AccessoryRarityUpResponse {
	userAccessory := GetUserAccessory(session, userAccessoryId)
	masterAccessory := session.Gamedata.Accessory[userAccessory.AccessoryMasterId]
	masterAfterAccessory := masterAccessory.RarityUp.AfterAccessory

	resp := response.AccessoryRarityUpResponse{
		DoRarityUp: client.DoRarityUp{
			BeforeAccessoryRarity: masterAccessory.Rarity.RarityType,
			AfterAccessoryRarity:  masterAfterAccessory.RarityType,
			DoRarityUpAddSkill:    userAccessory.PassiveSkill1Id.Value != *masterAfterAccessory.Grade[0].PassiveSkill1MasterId,
		},
		UserModelDiff: &session.UserModel,
	}

	// update the accessory
	userAccessory.AccessoryMasterId = masterAfterAccessory.Id
	userAccessory.Level = 1
	userAccessory.Exp = 0
	userAccessory.Grade = 0
	userAccessory.PassiveSkill1Id = generic.NewNullable(*masterAfterAccessory.Grade[0].PassiveSkill1MasterId)
	if masterAfterAccessory.Grade[0].PassiveSkill2MasterId != nil {
		userAccessory.PassiveSkill2Id = generic.NewNullable(*masterAfterAccessory.Grade[0].PassiveSkill2MasterId)
	}
	userAccessory.AcquiredAt = session.Time.Unix()
	UpdateUserAccessory(session, userAccessory)
	// remove resource used
	user_content.RemoveContent(session, masterAccessory.RarityUp.RarityUpGroup.Resource)
	user_content.RemoveContent(session, item.Gold.Amount(masterAccessory.Rarity.RarityUpMoney))
	return resp
}