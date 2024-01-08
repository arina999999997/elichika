package handler

import (
	"elichika/client"
	"elichika/config"
	"elichika/gamedata"
	"elichika/generic"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func AccessoryUpdateIsLock(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type UpdateIsLockReq struct {
		UserAccessoryId int64 `xorm:"'user_accessory_id' pk" json:"user_accessory_id"`
		IsLock          bool  `xorm:"'is_lock'" json:"is_lock"`
	}
	req := UpdateIsLockReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	accessory := session.GetUserAccessory(req.UserAccessoryId)
	accessory.IsLock = req.IsLock
	session.UpdateUserAccessory(accessory)

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AccessoryUpdateIsNew(ctx *gin.Context) {
	// this has no body or response, we just need to update every new accessory as not so
	// this can probably be optimised to a single SQL but no need to be so far
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	accessories := session.GetAllUserAccessories()
	for _, accessory := range accessories {
		accessory.IsNew = false
		session.UpdateUserAccessory(accessory)
	}
	session.Finalize("{}", "dummy")
	resp := SignResp(ctx, "{}", config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AccessoryMelt(ctx *gin.Context) {
	// disassemble
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type MeltReq struct {
		UserAccessoryIds []int64 `json:"user_accessory_ids"`
	}
	req := MeltReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	for _, userAccessoryId := range req.UserAccessoryIds {
		accessory := session.GetUserAccessory(userAccessoryId)
		session.AddResource(gamedata.Accessory[accessory.AccessoryMasterId].MeltGroup[accessory.Grade].Reward)
		accessory.IsNull = true // marked for delete
		session.UpdateUserAccessory(accessory)
	}

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AccessoryPowerUp(ctx *gin.Context) {
	// accessory synthesize
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()

	type PowerUpReq struct {
		UserAccessoryId int64 `json:"user_accessory_id"`
		// TODO: do not use anon type
		PowerUpAccessoryIds   []int64 `json:"power_up_user_accessory_ids"`
		AccessoryLevelUpItems []struct {
			AccessoryLevelUpItemMasterId int32 `json:"accessory_level_up_item_master_id"`
			Amount                       int32 `json:"amount"`
		} `json:"accessory_level_up_items"`
	}
	req := PowerUpReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	// limit break (grade up) is processed first, then exp is processed later
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	userAccessory := session.GetUserAccessory(req.UserAccessoryId)
	masterAccessory := gamedata.Accessory[userAccessory.AccessoryMasterId]

	skillPlusPercent := int32(0)
	moneyUsed := int32(0)

	type AccessoryDoPowerUp struct {
		DoLevelUp        bool `json:"do_level_up"`
		DoGradeUp        bool `json:"do_grade_up"`
		DoAddSkill       bool `json:"do_add_skill"`
		DoSkillProcessed bool `json:"do_skill_processed"`
		DoSkillLevelUp   bool `json:"do_skill_level_up"`
	}
	doPowerUp := AccessoryDoPowerUp{}
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
			doPowerUp.DoGradeUp = true
		} else {
			userAccessory.Exp += masterPowerUpAccessory.Rarity.LevelUp[powerUpAccessory.Level].PlusExp
			moneyUsed += masterPowerUpAccessory.Rarity.LevelUp[powerUpAccessory.Level].GameMoney
			skillPlusPercent += masterPowerUpAccessory.Rarity.SkillLevelUpPlusPercent[powerUpAccessory.PassiveSkill1Level.Value]
		}
		powerUpAccessory.IsNull = true // mark for delete
		session.UpdateUserAccessory(powerUpAccessory)
	}

	for _, item := range req.AccessoryLevelUpItems {
		itemId := item.AccessoryLevelUpItemMasterId
		session.RemoveResource(client.Content{
			ContentType:   24,
			ContentId:     int32(itemId),
			ContentAmount: int32(item.Amount),
		})
		userAccessory.Exp += item.Amount * gamedata.AccessoryLevelUpItem[itemId].PlusExp
		moneyUsed += item.Amount * gamedata.AccessoryLevelUpItem[itemId].GameMoney
	}

	// calculate new level
	// O(n) for now but can easily be O(log(n))
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
			doPowerUp.DoSkillLevelUp = true
		}
		doPowerUp.DoSkillProcessed = true
	}
	session.UpdateUserAccessory(userAccessory)

	session.RemoveGameMoney(int32(moneyUsed))

	signBody := session.Finalize("{}", "user_model_diff")
	// not sure what this is, doesn't seem to do anything
	// 10/20/30 but that's doesn't seems to be tied to rarity
	// setting it up to rarity doesn't hurt

	signBody, _ = sjson.Set(signBody, "success", masterAccessory.RarityType)
	signBody, _ = sjson.Set(signBody, "do_power_up", doPowerUp)

	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AccessoryRarityUp(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type RarityUpReq struct {
		UserAccessoryId int64 `xorm:"'user_accessory_id' pk" json:"user_accessory_id"`
	}
	req := RarityUpReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	userAccessory := session.GetUserAccessory(req.UserAccessoryId)
	masterAccessory := gamedata.Accessory[userAccessory.AccessoryMasterId]

	masterAfterAccessory := masterAccessory.RarityUp.AfterAccessory

	type AccessoryDoRarityUp struct {
		BeforeAccessoryRarity int32 `json:"before_accessory_rarity"`
		AfterAccessoryRarity  int32 `json:"after_accessory_rarity"`
		DoRarityUpAddSkill    bool  `json:"do_rarity_up_add_skill"`
	}
	doRarityUp := AccessoryDoRarityUp{
		BeforeAccessoryRarity: masterAccessory.RarityType,
		AfterAccessoryRarity:  masterAfterAccessory.RarityType,
		DoRarityUpAddSkill:    false,
	}
	// update the accessory
	userAccessory.AccessoryMasterId = masterAfterAccessory.Id
	userAccessory.Level = 1
	userAccessory.Exp = 0
	userAccessory.Grade = 0
	doRarityUp.DoRarityUpAddSkill = userAccessory.PassiveSkill1Id.Value != *masterAfterAccessory.Grade[0].PassiveSkill1MasterId
	userAccessory.PassiveSkill1Id = generic.NewNullable(*masterAfterAccessory.Grade[0].PassiveSkill1MasterId)
	if masterAfterAccessory.Grade[0].PassiveSkill2MasterId != nil {
		userAccessory.PassiveSkill2Id = generic.NewNullable(*masterAfterAccessory.Grade[0].PassiveSkill2MasterId)
		// userAccessory.PassiveSkill2Id = new(int)
		// *userAccessory.PassiveSkill2Id = *masterAfterAccessory.Grade[0].PassiveSkill2MasterId
	}
	userAccessory.AcquiredAt = session.Time.Unix()
	session.UpdateUserAccessory(userAccessory)
	// remove resource used
	session.RemoveResource(masterAccessory.RarityUp.RarityUpGroup.Resource)
	session.RemoveGameMoney(int32(masterAccessory.Rarity.RarityUpMoney))

	// finalize and send the response

	signBody := session.Finalize("{}", "user_model_diff")
	signBody, _ = sjson.Set(signBody, "do_rarity_up", doRarityUp)
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AccessoryAllUnequip(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type AllUnequipReq struct {
		UserAccessoryId int64 `xorm:"'user_accessory_id' pk" json:"user_accessory_id"`
	}
	req := AllUnequipReq{}
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
	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
