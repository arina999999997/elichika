package user_accessory

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/item"
	"elichika/subsystem/user_content"
	"elichika/userdata"

	"math/rand"
)

func LevelUpAccessory(session *userdata.Session, userAccessoryId int64,
	powerUpAccessoryIds generic.Array[int64],
	accessoryLevelUpItems generic.Array[client.AccessoryLevelUpItem]) response.AccessoryPowerUpResponse {

	userAccessory := GetUserAccessory(session, userAccessoryId)
	masterAccessory := session.Gamedata.Accessory[userAccessory.AccessoryMasterId]

	resp := response.AccessoryPowerUpResponse{
		UserModelDiff: &session.UserModel,
	}

	expGain := int32(0)
	skillPlusPercent := int32(0)
	moneyUsed := int32(0)

	// power up is processed by listing order
	// so different order of accessory can result in different result
	for _, powerUpAccessoryId := range powerUpAccessoryIds.Slice {
		powerUpAccessory := GetUserAccessory(session, powerUpAccessoryId)
		masterPowerUpAccessory := session.Gamedata.Accessory[powerUpAccessory.AccessoryMasterId]

		if (userAccessory.Grade < 5) && (powerUpAccessory.AccessoryMasterId == userAccessory.AccessoryMasterId) {
			// limit increase
			userAccessory.Grade += powerUpAccessory.Grade + 1
			if userAccessory.Grade > 5 {
				userAccessory.Grade = 5
			}
			moneyUsed += masterPowerUpAccessory.Rarity.GradeUpMoney[powerUpAccessory.Grade]

			// some limit increase change the skills
			if masterAccessory.Grade[userAccessory.Grade].PassiveSkill1MasterId != nil {
				userAccessory.PassiveSkill1Id = generic.NewNullable(*masterAccessory.Grade[userAccessory.Grade].PassiveSkill1MasterId)
			}
			if masterAccessory.Grade[userAccessory.Grade].PassiveSkill2MasterId != nil {
				userAccessory.PassiveSkill2Id = generic.NewNullable(*masterAccessory.Grade[userAccessory.Grade].PassiveSkill2MasterId)
			}
			resp.DoPowerUp.DoGradeUp = true
		} else {
			expGain += masterPowerUpAccessory.Rarity.LevelUp[powerUpAccessory.Level].PlusExp
			moneyUsed += masterPowerUpAccessory.Rarity.LevelUp[powerUpAccessory.Level].GameMoney
			skillPlusPercent += masterPowerUpAccessory.Rarity.SkillLevelUpPlusPercent[powerUpAccessory.PassiveSkill1Level.Value]
		}
		DeleteUserAccessory(session, powerUpAccessory.UserAccessoryId)
	}

	for _, item := range accessoryLevelUpItems.Slice {
		itemId := item.AccessoryLevelUpItemMasterId
		user_content.RemoveContent(session, client.Content{
			ContentType:   enum.ContentTypeAccessoryLevelUp,
			ContentId:     itemId,
			ContentAmount: item.Amount,
		})
		expGain += item.Amount * session.Gamedata.AccessoryLevelUpItem[itemId].PlusExp
		moneyUsed += item.Amount * session.Gamedata.AccessoryLevelUpItem[itemId].GameMoney
	}

	{
		// the bonus is correct but who knows what the chance actually are
		successType := rand.Intn(3)
		switch successType {
		case 0:
			resp.Success = enum.AccessoryLevelUpSuccessNormalSuccess
		case 1:
			resp.Success = enum.AccessoryLevelUpSuccessBigSuccess
			expGain = expGain * 3 / 2
		case 2:
			resp.Success = enum.AccessoryLevelUpSuccessSuperSuccess
			expGain = expGain * 2
		}
	}

	// calculate new level
	// O(n) for now but can easily be O(log(n))
	userAccessory.Exp += expGain
	for userAccessory.Level < masterAccessory.Grade[userAccessory.Grade].MaxLevel {
		if masterAccessory.LevelExp[userAccessory.Level+1] <= userAccessory.Exp {
			userAccessory.Level++
		} else {
			break
		}
	}

	// remove extra exp
	if userAccessory.Level == masterAccessory.Grade[userAccessory.Grade].MaxLevel {
		userAccessory.Exp = masterAccessory.LevelExp[userAccessory.Level]
	}

	// calculate new skill level
	if userAccessory.PassiveSkill1Level.Value < masterAccessory.Rarity.GradeMaxSkillLevel[userAccessory.Grade] {
		denominator := masterAccessory.Rarity.SkillLevelUpDenominator[userAccessory.Grade][userAccessory.PassiveSkill1Level.Value]
		chance := skillPlusPercent / denominator
		// probability is chance / 10000
		if int32(rand.Intn(10000)) < chance {
			userAccessory.PassiveSkill1Level.Value++
			resp.DoPowerUp.DoSkillLevelUp = true
		}
		resp.DoPowerUp.DoSkillProcessed = true
	}
	UpdateUserAccessory(session, userAccessory)
	user_content.RemoveContent(session, item.Gold.Amount(moneyUsed))
	return resp
}
