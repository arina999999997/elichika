package handler

import (
	"elichika/config"
	"elichika/gamedata"
	"elichika/model"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func AccessoryUpdateIsLock(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type UpdateIsLockReq struct {
		UserAccessoryID int64 `xorm:"'user_accessory_id' pk" json:"user_accessory_id"`
		IsLock          bool  `xorm:"'is_lock'" json:"is_lock"`
	}
	req := UpdateIsLockReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	accessory := session.GetUserAccessory(req.UserAccessoryID)
	accessory.IsLock = req.IsLock
	session.UpdateUserAccessory(accessory)

	signBody := session.Finalize(GetData("userModel.json"), "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AccessoryUpdateIsNew(ctx *gin.Context) {
	// this has no body or response, we just need to update every new accessory as not so
	// this can probably be optimised to a single SQL but no need to be so far
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	accessories := session.GetAllUserAccessories()
	for _, accessory := range accessories {
		accessory.IsNew = false
		session.UpdateUserAccessory(accessory)
	}
	session.Finalize("", "")
	resp := SignResp(ctx, "{}", config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AccessoryMelt(ctx *gin.Context) {
	// disassemble
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type MeltReq struct {
		UserAccessoryIDs []int64 `json:"user_accessory_ids"`
	}
	req := MeltReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	for _, userAccessoryID := range req.UserAccessoryIDs {
		accessory := session.GetUserAccessory(userAccessoryID)
		meltGroupID := gamedata.Accessory.Accessory[accessory.AccessoryMasterID].Grade[accessory.Grade].MeltGroupMasterID
		session.AddResource(gamedata.Accessory.MeltGroup[meltGroupID].Resource)
		accessory.AccessoryMasterID = 0 // marked for delete
		session.UpdateUserAccessory(accessory)
	}

	signBody := session.Finalize(GetData("userModel.json"), "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AccessoryPowerUp(ctx *gin.Context) {
	// accessory synthesize
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()

	type PowerUpReq struct {
		UserAccessoryID       int64   `json:"user_accessory_id"`
		PowerUpAccessoryIDs   []int64 `json:"power_up_user_accessory_ids"`
		AccessoryLevelUpItems []struct {
			AccessoryLevelUpItemMasterID int `json:"accessory_level_up_item_master_id"`
			Amount                       int `json:"amount"`
		} `json:"accessory_level_up_items"`
	}
	req := PowerUpReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	// limit break (grade up) is processed first, then exp is processed later
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	userAccessory := session.GetUserAccessory(req.UserAccessoryID)
	masterAccessory := gamedata.Accessory.Accessory[userAccessory.AccessoryMasterID]

	skillPlusPercent := 0
	moneyUsed := 0

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
	for _, powerUpAccessoryID := range req.PowerUpAccessoryIDs {
		powerUpAccessory := session.GetUserAccessory(powerUpAccessoryID)
		powerUpRarity := gamedata.Accessory.Accessory[powerUpAccessory.AccessoryMasterID].RarityType

		if (userAccessory.Grade < 5) && (powerUpAccessory.AccessoryMasterID == userAccessory.AccessoryMasterID) {
			// limit increase
			userAccessory.Grade += powerUpAccessory.Grade + 1
			if userAccessory.Grade > 5 {
				userAccessory.Grade = 5
			}
			moneyUsed += gamedata.Accessory.Rarity[powerUpRarity].Grade[powerUpAccessory.Grade].GradeUpMoney

			// some limit increase change the skills
			if masterAccessory.Grade[userAccessory.Grade].PassiveSkill1MasterID != nil {
				userAccessory.PassiveSkill1ID = *masterAccessory.Grade[userAccessory.Grade].PassiveSkill1MasterID
			}
			if masterAccessory.Grade[userAccessory.Grade].PassiveSkill2MasterID != nil {
				userAccessory.PassiveSkill2ID = new(int)
				*userAccessory.PassiveSkill2ID = *masterAccessory.Grade[userAccessory.Grade].PassiveSkill2MasterID
			}
			doPowerUp.DoGradeUp = true
		} else {
			userAccessory.Exp += gamedata.Accessory.Rarity[powerUpRarity].Level[powerUpAccessory.Level].PlusExp
			moneyUsed += gamedata.Accessory.Rarity[powerUpRarity].Level[powerUpAccessory.Level].GameMoney
			skillPlusPercent += gamedata.Accessory.Rarity[powerUpRarity].SkillLevel[powerUpAccessory.PassiveSkill1Level].PlusPercent
		}
		powerUpAccessory.AccessoryMasterID = 0 // mark for delete
		session.UpdateUserAccessory(powerUpAccessory)
	}

	for _, item := range req.AccessoryLevelUpItems {
		itemID := item.AccessoryLevelUpItemMasterID
		session.RemoveResource(model.Content{
			ContentType:   24,
			ContentID:     itemID,
			ContentAmount: int64(item.Amount),
		})
		userAccessory.Exp += item.Amount * gamedata.Accessory.LevelUpItem[itemID].PlusExp
		moneyUsed += item.Amount * gamedata.Accessory.LevelUpItem[itemID].GameMoney
	}

	// calculate new level
	// O(n) for now but can easily be O(log(n))
	for userAccessory.Level < masterAccessory.Grade[userAccessory.Grade].MaxLevel {
		if masterAccessory.Level[userAccessory.Level+1].Exp <= userAccessory.Exp {
			userAccessory.Level++
		} else {
			break
		}
	}

	// remove extra exp
	if userAccessory.Level == masterAccessory.Grade[userAccessory.Grade].MaxLevel {
		userAccessory.Exp = masterAccessory.Level[userAccessory.Level].Exp
	}

	// calculate new skill level
	if userAccessory.PassiveSkill1Level < gamedata.Accessory.Rarity[masterAccessory.RarityType].Grade[userAccessory.Grade].SkillMaxLevel {
		denominator := gamedata.Accessory.Rarity[masterAccessory.RarityType].Grade[userAccessory.Grade].SkillLevelUpDenominator
		chance := skillPlusPercent / denominator[userAccessory.PassiveSkill1Level]
		// probability is chance / 10000
		if chance > rand.Intn(10000) {
			userAccessory.PassiveSkill1Level++
			doPowerUp.DoSkillLevelUp = true
		}
		doPowerUp.DoSkillProcessed = true
	}
	session.UpdateUserAccessory(userAccessory)

	session.RemoveGameMoney(int64(moneyUsed))

	signBody := session.Finalize(GetData("userModelDiff.json"), "user_model_diff")
	// not sure what this is, doesn't seem to do anything
	// 10/20/30 but that's doesn't seems to be tied to rarity
	// setting it up to rarity doesn't hurt

	signBody, _ = sjson.Set(signBody, "success", masterAccessory.RarityType)
	signBody, _ = sjson.Set(signBody, "do_power_up", doPowerUp)

	// fmt.Println("Money used: ", moneyUsed, "Skill plus percent: ", skillPlusPercent)
	// fmt.Println("New level: ", userAccessory.Level, "New exp: ", userAccessory.Exp)
	// fmt.Println(signBody)
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AccessoryRarityUp(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type RarityUpReq struct {
		UserAccessoryID int64 `xorm:"'user_accessory_id' pk" json:"user_accessory_id"`
	}
	req := RarityUpReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	userAccessory := session.GetUserAccessory(req.UserAccessoryID)
	masterAccessory := gamedata.Accessory.Accessory[userAccessory.AccessoryMasterID]

	masterAfterAccessory := gamedata.Accessory.Accessory[masterAccessory.RarityUp.AfterAccessoryMasterID]

	type AccessoryDoRarityUp struct {
		BeforeAccessoryRarity int  `json:"before_accessory_rarity"`
		AfterAccessoryRarity  int  `json:"after_accessory_rarity"`
		DoRarityUpAddSkill    bool `json:"do_rarity_up_add_skill"`
	}
	doRarityUp := AccessoryDoRarityUp{
		BeforeAccessoryRarity: masterAccessory.RarityType,
		AfterAccessoryRarity:  masterAfterAccessory.RarityType,
		DoRarityUpAddSkill:    false,
	}
	// update the accessory
	userAccessory.AccessoryMasterID = masterAfterAccessory.MasterID
	userAccessory.Level = 1
	userAccessory.Exp = 0
	userAccessory.Grade = 0
	doRarityUp.DoRarityUpAddSkill = userAccessory.PassiveSkill1ID != *masterAfterAccessory.Grade[0].PassiveSkill1MasterID
	userAccessory.PassiveSkill1ID = *masterAfterAccessory.Grade[0].PassiveSkill1MasterID
	if masterAfterAccessory.Grade[0].PassiveSkill2MasterID != nil {
		userAccessory.PassiveSkill2ID = new(int)
		*userAccessory.PassiveSkill2ID = *masterAfterAccessory.Grade[0].PassiveSkill2MasterID
	}
	userAccessory.AcquiredAt = time.Now().Unix()
	session.UpdateUserAccessory(userAccessory)
	// remove resource used
	session.RemoveResource(gamedata.Accessory.RarityUpGroup[masterAccessory.RarityUp.AccessoryRarityUpGroupMasterID].Resource)
	session.RemoveGameMoney(int64(gamedata.Accessory.Rarity[masterAccessory.RarityType].RarityUpMoney))

	// finalize and send the response

	signBody := session.Finalize(GetData("userModelDiff.json"), "user_model_diff")
	signBody, _ = sjson.Set(signBody, "do_rarity_up", doRarityUp)
	resp := SignResp(ctx, signBody, config.SessionKey)
	// fmt.Println(resp)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AccessoryAllUnequip(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type AllUnequipReq struct {
		UserAccessoryID int64 `xorm:"'user_accessory_id' pk" json:"user_accessory_id"`
	}
	req := AllUnequipReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	liveParties := session.GetAllLiveParties()
	for _, liveParty := range liveParties {
		updated := false
		if (liveParty.UserAccessoryID1 != nil) && (*liveParty.UserAccessoryID1 == req.UserAccessoryID) {
			liveParty.UserAccessoryID1 = nil
			updated = true
		}
		if (liveParty.UserAccessoryID2 != nil) && (*liveParty.UserAccessoryID2 == req.UserAccessoryID) {
			liveParty.UserAccessoryID2 = nil
			updated = true
		}
		if (liveParty.UserAccessoryID3 != nil) && (*liveParty.UserAccessoryID3 == req.UserAccessoryID) {
			liveParty.UserAccessoryID3 = nil
			updated = true
		}
		if updated {
			session.UpdateUserLiveParty(liveParty)
		}
	}
	signBody := session.Finalize(GetData("userModel.json"), "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	// fmt.Println(resp)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
