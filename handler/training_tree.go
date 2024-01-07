package handler

import (
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/model"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func FetchTrainingTree(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type FetchTrainingTreeReq struct {
		CardMasterId int `json:"card_master_id"`
	}
	req := FetchTrainingTreeReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	signBody := `"{}"`
	signBody, _ = sjson.Set(signBody, "user_card_training_tree_cell_list", session.GetTrainingTree(req.CardMasterId))
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func LevelUpCard(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseTrainingLevelUp {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseTrainingActivateCell
	}

	type LevelUpCardReq struct {
		CardMasterId    int `json:"card_master_id"`
		AdditionalLevel int `json:"additional_level"`
	}

	req := LevelUpCardReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	cardLevel := gamedata.CardLevel[gamedata.Card[req.CardMasterId].CardRarityType]
	card := session.GetUserCard(req.CardMasterId)
	session.RemoveGameMoney(int32(
		cardLevel.GameMoneyPrefixSum[card.Level+req.AdditionalLevel] - cardLevel.GameMoneyPrefixSum[card.Level]))
	session.RemoveCardExp(int32(
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
		CardMasterId int `json:"card_master_id"`
		ContentId    int `json:"content_id"`
	}
	req := GradeUpCardReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)
	masterCard := gamedata.Card[req.CardMasterId]
	card := session.GetUserCard(req.CardMasterId)
	member := session.GetMember(*masterCard.MemberMasterId)
	card.Grade++
	currentLoveLevel := gamedata.LoveLevelFromLovePoint(member.LovePointLimit)
	currentLoveLevel += masterCard.CardRarityType / 10 // TODO: Do not hard code this
	if currentLoveLevel > gamedata.MemberLoveLevelCount {
		currentLoveLevel = gamedata.MemberLoveLevelCount
	}
	member.LovePointLimit = gamedata.MemberLoveLevelLovePoint[currentLoveLevel]
	session.UpdateUserCard(card)
	member.IsNew = true
	session.UpdateMember(member)
	session.RemoveResource(masterCard.CardGradeUpItem[card.Grade][int32(req.ContentId)])
	// we need to set user_info_trigger_card_grade_up_by_trigger_id
	// for the pop up after limit breaking
	// this trigger show the pop up after limit break
	session.AddTriggerCardGradeUp(model.TriggerCardGradeUp{
		CardMasterId:         req.CardMasterId,
		BeforeLoveLevelLimit: currentLoveLevel - masterCard.CardRarityType/10,
		AfterLoveLevelLimit:  currentLoveLevel})

	resp := session.Finalize("{}", "user_model_diff")
	resp = SignResp(ctx, resp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func ActivateTrainingTreeCell(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type ActivateTrainingTreeCellReq struct {
		CardMasterId  int   `json:"card_master_id"`
		CellMasterIds []int `json:"cell_master_ids"`
		PayType       int   `json:"pay_type"`
	}
	req := ActivateTrainingTreeCellReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseTrainingActivateCell {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseDeckEdit
	}

	card := session.GetUserCard(req.CardMasterId)
	masterCard := gamedata.Card[req.CardMasterId]
	trainingTree := masterCard.TrainingTree
	cellContents := &trainingTree.TrainingTreeMapping.TrainingTreeCellContents
	for _, cellId := range req.CellMasterIds {
		cell := &(*cellContents)[cellId]
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
			naviActionId := trainingTree.NaviActionIds[cell.TrainingContentNo]
			session.UpdateVoice(naviActionId, true)
		case 4: // story cell
			// training_content_type 11 in m_training_tree_card_story_side
			storySideId, exist := trainingTree.TrainingTreeCardStorySides[11]
			if !exist {
				panic("story doesn't exist")
			}
			session.InsertStorySide(storySideId)
		case 5:
			// idolize
			card.IsAwakening = true
			card.IsAwakeningImage = true
			storySideId, exist := trainingTree.TrainingTreeCardStorySides[9]
			if exist {
				session.InsertStorySide(storySideId)
			}
		case 6: // costume
			// alternative suit is awarded based on amount of tile instead
			session.InsertUserSuit(trainingTree.SuitMIds[cell.TrainingContentNo])
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
		if reward.ActivateNum > card.TrainingActivatedCellCount+len(req.CellMasterIds) {
			break
		}
		if reward.ActivateNum > card.TrainingActivatedCellCount {
			session.AddResource(reward.Reward)
		}
	}

	card.TrainingActivatedCellCount += len(req.CellMasterIds)

	if card.TrainingActivatedCellCount+1 == len(*cellContents) {
		card.IsAllTrainingActivated = true
		member := session.GetMember(*masterCard.MemberMasterId)
		member.AllTrainingCardCount++
		session.UpdateMember(member)
	}

	session.UpdateUserCard(card)

	// set "user_card_training_tree_cell_list" to the cell unlocked and insert the cell to db
	unlockedCells := []model.TrainingTreeCell{}
	for _, cellId := range req.CellMasterIds {
		unlockedCells = append(unlockedCells,
			model.TrainingTreeCell{
				UserId:       userId,
				CardMasterId: req.CardMasterId,
				CellId:       cellId,
				ActivatedAt:  session.Time.Unix()})
	}

	session.InsertTrainingTreeCells(unlockedCells)

	jsonResp := session.Finalize("{}", "user_model_diff")
	jsonResp, _ = sjson.Set(jsonResp, "user_card_training_tree_cell_list", session.GetTrainingTree(req.CardMasterId))
	resp := SignResp(ctx, jsonResp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
