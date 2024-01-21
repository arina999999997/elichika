package accessory

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func AccessoryUpdateIsLock(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.AccessoryUpdateIsLockRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	accessory := session.GetUserAccessory(req.UserAccessoryId)
	accessory.IsLock = req.IsLock
	session.UpdateUserAccessory(accessory)

	session.Finalize()
	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func AccessoryUpdateIsNew(ctx *gin.Context) {
	// this has no body or response, we just need to update every new accessory as not so
	// this can be optimised to a single sql call but we don't need to take it that far (yet)
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	accessories := session.GetAllUserAccessories()
	for _, accessory := range accessories {
		accessory.IsNew = false
		session.UpdateUserAccessory(accessory)
	}
	session.Finalize()
	common.JsonResponse(ctx, &response.EmptyResponse{})
}

func AccessoryMelt(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.AccessoryMeltRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	for _, userAccessoryId := range req.UserAccessoryIds {
		accessory := session.GetUserAccessory(userAccessoryId)
		session.AddResource(gamedata.Accessory[accessory.AccessoryMasterId].MeltGroup[accessory.Grade].Reward)
		session.DeleteUserAccessory(userAccessoryId)
	}

	session.Finalize()
	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func AccessoryPowerUp(ctx *gin.Context) {
	// accessory synthesize
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.AccessoryPowerUpRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	// limit break (grade up) is processed first, then exp is processed later
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	userAccessory := session.GetUserAccessory(req.UserAccessoryId)
	masterAccessory := gamedata.Accessory[userAccessory.AccessoryMasterId]

	resp := response.AccessoryPowerUpResponse{
		UserModelDiff: &session.UserModel,
	}

	expGain := int32(0)
	skillPlusPercent := int32(0)
	moneyUsed := int32(0)

	// power up is processed by listing order
	// so different order of accessory can result in different result
	for _, powerUpAccessoryId := range req.PowerUpAccessoryIds {
		powerUpAccessory := session.GetUserAccessory(powerUpAccessoryId)
		masterPowerUpAccessory := gamedata.Accessory[powerUpAccessory.AccessoryMasterId]

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
		session.DeleteUserAccessory(powerUpAccessory.UserAccessoryId)
	}

	for _, item := range req.AccessoryLevelUpItems {
		itemId := item.AccessoryLevelUpItemMasterId
		// TODO: maybe make the item into a map at the start?
		session.RemoveResource(client.Content{
			ContentType:   enum.ContentTypeAccessoryLevelUp,
			ContentId:     itemId,
			ContentAmount: item.Amount,
		})
		expGain += item.Amount * gamedata.AccessoryLevelUpItem[itemId].PlusExp
		moneyUsed += item.Amount * gamedata.AccessoryLevelUpItem[itemId].GameMoney
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
	session.UpdateUserAccessory(userAccessory)
	session.RemoveGameMoney(int32(moneyUsed))

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func AccessoryRarityUp(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.AccessoryRarityUpRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	userAccessory := session.GetUserAccessory(req.UserAccessoryId)
	masterAccessory := gamedata.Accessory[userAccessory.AccessoryMasterId]
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
	session.UpdateUserAccessory(userAccessory)
	// remove resource used
	session.RemoveResource(masterAccessory.RarityUp.RarityUpGroup.Resource)
	session.RemoveGameMoney(int32(masterAccessory.Rarity.RarityUpMoney))

	// finalize and send the response

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func AccessoryAllUnequip(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.AccessoryAllUnequipRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	liveParties := session.GetAllLivePartiesWithAccessory(req.UserAccessoryId)

	for _, liveParty := range liveParties {
		if (liveParty.UserAccessoryId1.HasValue) && (liveParty.UserAccessoryId1.Value == req.UserAccessoryId) {
			liveParty.UserAccessoryId1.HasValue = false
		}
		if (liveParty.UserAccessoryId2.HasValue) && (liveParty.UserAccessoryId2.Value == req.UserAccessoryId) {
			liveParty.UserAccessoryId2.HasValue = false
		}
		if (liveParty.UserAccessoryId3.HasValue) && (liveParty.UserAccessoryId3.Value == req.UserAccessoryId) {
			liveParty.UserAccessoryId3.HasValue = false
		}
		session.UpdateUserLiveParty(liveParty)
	}
	session.Finalize()

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/accessory/updateIsLock", AccessoryUpdateIsLock)
	router.AddHandler("/accessory/updateIsNew", AccessoryUpdateIsNew)
	router.AddHandler("/accessory/melt", AccessoryMelt)
	router.AddHandler("/accessory/powerUp", AccessoryPowerUp)
	router.AddHandler("/accessory/rarityUp", AccessoryRarityUp)
	router.AddHandler("/accessory/allUnequip", AccessoryAllUnequip)
}
