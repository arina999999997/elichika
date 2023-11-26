package handler

import (
	"elichika/config"
	"elichika/gamedata"
	"elichika/model"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func FetchTrainingTree(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type FetchTrainingTreeReq struct {
		CardMasterID int `json:"card_master_id"`
	}
	req := FetchTrainingTreeReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	signBody := `"{}"`
	signBody, _ = sjson.Set(signBody, "user_card_training_tree_cell_list", session.GetTrainingTree(req.CardMasterID))
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func LevelUpCard(ctx *gin.Context) {
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()

	type LevelUpCardReq struct {
		CardMasterID    int `json:"card_master_id"`
		AdditionalLevel int `json:"additional_level"`
	}

	req := LevelUpCardReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	cardLevel := gamedata.CardLevel[gamedata.Card[req.CardMasterID].CardRarityType]
	card := session.GetUserCard(req.CardMasterID)
	session.RemoveGameMoney(int64(
		cardLevel.GameMoneyPrefixSum[card.Level+req.AdditionalLevel] - cardLevel.GameMoneyPrefixSum[card.Level]))
	session.RemoveCardExp(int64(
		cardLevel.ExpPrefixSum[card.Level+req.AdditionalLevel] - cardLevel.ExpPrefixSum[card.Level]))
	card.Level += req.AdditionalLevel
	session.UpdateUserCard(card)
	signBody := session.Finalize("{}", "user_model_diff")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func GradeUpCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type GradeUpCardReq struct {
		CardMasterID int `json:"card_master_id"`
		ContentID    int `json:"content_id"`
	}
	req := GradeUpCardReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	masterCard := gamedata.Card[req.CardMasterID]
	card := session.GetUserCard(req.CardMasterID)
	member := session.GetMember(*masterCard.MemberMasterID)
	card.Grade++
	currentBondLevel := gamedata.LoveLevelFromLovePoint(member.LovePointLimit)
	currentBondLevel += masterCard.CardRarityType / 10 // TODO: Do not hard code this
	if currentBondLevel > gamedata.MemberLoveLevelCount {
		currentBondLevel = gamedata.MemberLoveLevelCount
	}
	member.LovePointLimit = gamedata.MemberLoveLevelLovePoint[currentBondLevel]
	session.UpdateUserCard(card)
	member.IsNew = true
	session.UpdateMember(member)
	session.RemoveResource(masterCard.CardGradeUpItem[card.Grade][req.ContentID])
	// we need to set user_info_trigger_card_grade_up_by_trigger_id
	// for the pop up after limit breaking
	// this trigger show the pop up after limit break
	session.AddTriggerCardGradeUp(model.TriggerCardGradeUp{
		CardMasterID:         req.CardMasterID,
		BeforeLoveLevelLimit: currentBondLevel - masterCard.CardRarityType/10,
		AfterLoveLevelLimit:  currentBondLevel})

	resp := session.Finalize("{}", "user_model_diff")
	resp = SignResp(ctx, resp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ActivateTrainingTreeCell(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type ActivateTrainingTreeCellReq struct {
		CardMasterID  int   `json:"card_master_id"`
		CellMasterIDs []int `json:"cell_master_ids"`
		PayType       int   `json:"pay_type"`
	}
	req := ActivateTrainingTreeCellReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	card := session.GetUserCard(req.CardMasterID)
	masterCard := gamedata.Card[req.CardMasterID]
	trainingTree := masterCard.TrainingTree
	cellContents := &trainingTree.TrainingTreeMapping.TrainingTreeCellContents
	for _, cellID := range req.CellMasterIDs {
		cell := &(*cellContents)[cellID]
		// consume practice items
		for _, resource := range cell.TrainingTreeCellItemSet.Resources {
			session.RemoveResource(resource)
		}

		switch cell.TrainingTreeCellType {
		// note that this reference is different from m_training_tree_cell_type_setting
		// that table is for training_content_type
		case 2: // params
			paramCell := &trainingTree.TrainingTreeCardParams[cell.TrainingContentNo]
			switch paramCell.TrainingContentType {
			case 2: // stamina
				card.TrainingLife += paramCell.Value
			case 3: // appeal
				card.TrainingAttack += paramCell.Value
			case 4: // technique
				card.TrainingDexterity += paramCell.Value
			default:
				panic("Unexpected training content type")
			}
		case 3: // voice
			naviActionID := trainingTree.NaviActionIDs[cell.TrainingContentNo]
			session.UpdateVoice(naviActionID, true)
		case 4: // story cell
			// training_content_type 11 in m_training_tree_card_story_side
			storySideID, exist := trainingTree.TrainingTreeCardStorySides[11]
			if !exist {
				panic("story doesn't exist")
			}
			session.InsertStorySide(storySideID)
		case 5:
			// idolize
			card.IsAwakening = true
			card.IsAwakeningImage = true
			storySideID, exist := trainingTree.TrainingTreeCardStorySides[9]
			if exist {
				session.InsertStorySide(storySideID)
			}
		case 6: // costume
			// alternative suit is awarded based on amount of tile instead
			session.InsertUserSuit(trainingTree.SuitMIDs[cell.TrainingContentNo])
		case 7: // skill
			card.ActiveSkillLevel++
		case 8: // insight
			card.MaxFreePassiveSkill++
		case 9: // ability
			card.PassiveSkillALevel++
		default:
			panic("Unknown cell type")
		}
	}

	// progress reward
	for _, reward := range trainingTree.TrainingTreeProgressRewards {
		if reward.ActivateNum > card.TrainingActivatedCellCount+len(req.CellMasterIDs) {
			break
		}
		if reward.ActivateNum > card.TrainingActivatedCellCount {
			session.AddResource(reward.Reward)
		}
	}

	card.TrainingActivatedCellCount += len(req.CellMasterIDs)

	if card.TrainingActivatedCellCount+1 == len(*cellContents) {
		card.IsAllTrainingActivated = true
		member := session.GetMember(*masterCard.MemberMasterID)
		member.AllTrainingCardCount++
		session.UpdateMember(member)
	}

	session.UpdateUserCard(card)

	// set "user_card_training_tree_cell_list" to the cell unlocked and insert the cell to db
	unlockedCells := []model.TrainingTreeCell{}
	timeStamp := time.Now().Unix()
	for _, cellID := range req.CellMasterIDs {
		unlockedCells = append(unlockedCells,
			model.TrainingTreeCell{
				UserID:       userID,
				CardMasterID: req.CardMasterID,
				CellID:       cellID,
				ActivatedAt:  timeStamp})
	}

	session.InsertTrainingTreeCells(unlockedCells)

	jsonResp := session.Finalize("{}", "user_model_diff")
	jsonResp, _ = sjson.Set(jsonResp, "user_card_training_tree_cell_list", session.GetTrainingTree(req.CardMasterID))
	resp := SignResp(ctx, jsonResp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
