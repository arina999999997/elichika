package training_tree

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/item"
	"elichika/router"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_member"
	"elichika/subsystem/user_suit"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchTrainingTree(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchTrainingTreeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	common.JsonResponse(ctx, response.FetchTrainingTreeResponse{
		UserCardTrainingTreeCellList: session.GetTrainingTree(req.CardMasterId),
	})
}

func LevelUpCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.LevelUpCardRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseTrainingLevelUp {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseTrainingActivateCell
	}

	cardLevel := session.Gamedata.CardLevel[session.Gamedata.Card[req.CardMasterId].CardRarityType]
	card := session.GetUserCard(req.CardMasterId)
	user_content.RemoveContent(session, item.Gold.Amount(int32(
		cardLevel.GameMoneyPrefixSum[card.Level+req.AdditionalLevel]-cardLevel.GameMoneyPrefixSum[card.Level])))
	user_content.RemoveContent(session, item.EXP.Amount(int32(
		cardLevel.ExpPrefixSum[card.Level+req.AdditionalLevel]-cardLevel.ExpPrefixSum[card.Level])))
	card.Level += req.AdditionalLevel
	session.UpdateUserCard(card)

	session.Finalize()
	common.JsonResponse(ctx, response.LevelUpCardResponse{
		UserModelDiff: &session.UserModel,
	})
}

func GradeUpCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.GradeUpCardRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	masterCard := session.Gamedata.Card[req.CardMasterId]
	card := session.GetUserCard(req.CardMasterId)
	member := user_member.GetMember(session, *masterCard.MemberMasterId)

	card.Grade++
	currentLoveLevel := session.Gamedata.LoveLevelFromLovePoint(member.LovePointLimit)
	currentLoveLevel += masterCard.CardRarityType / 10 // TODO: Do not hard code this

	if currentLoveLevel > session.Gamedata.MemberLoveLevelCount {
		currentLoveLevel = session.Gamedata.MemberLoveLevelCount
	}
	member.LovePointLimit = session.Gamedata.MemberLoveLevelLovePoint[currentLoveLevel]
	session.UpdateUserCard(card)
	member.IsNew = true
	user_member.UpdateMember(session, member)
	user_content.RemoveContent(session, masterCard.CardGradeUpItem[card.Grade][req.ContentId])
	// we need to set user_info_trigger_card_grade_up_by_trigger_id
	// for the pop up after limit breaking
	// this trigger show the pop up after limit break
	session.AddTriggerCardGradeUp(client.UserInfoTriggerCardGradeUp{
		CardMasterId:         req.CardMasterId,
		BeforeLoveLevelLimit: int32(currentLoveLevel - masterCard.CardRarityType/10),
		AfterLoveLevelLimit:  int32(currentLoveLevel)})

	session.Finalize()
	common.JsonResponse(ctx, response.GradeUpCardResponse{
		UserModelDiff: &session.UserModel,
	})
}

func ActivateTrainingTreeCell(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ActivateTrainingTreeCellRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseTrainingActivateCell {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseDeckEdit
	}

	card := session.GetUserCard(req.CardMasterId)
	masterCard := session.Gamedata.Card[req.CardMasterId]
	trainingTree := masterCard.TrainingTree
	cellContents := trainingTree.TrainingTreeMapping.TrainingTreeCellContents
	for _, cellId := range req.CellMasterIds.Slice {
		cell := &cellContents[cellId]
		// consume practice items
		for _, resource := range cell.TrainingTreeCellItemSet.Resources {
			user_content.RemoveContent(session, resource)
		}

		switch int32(cell.TrainingTreeCellType) {
		// note that this reference is different from m_training_tree_cell_type_setting
		// that table is for training_content_type
		case enum.TrainingTreeCellTypeParameter:
			paramCell := &trainingTree.TrainingTreeCardParams[cell.TrainingContentNo]
			switch int32(paramCell.TrainingContentType) {
			case enum.TrainingContentTypeStamina: // stamina
				card.TrainingLife += int32(paramCell.Value)
			case enum.TrainingContentTypeAppeal: // appeal
				card.TrainingAttack += int32(paramCell.Value)
			case enum.TrainingContentTypeTechnique: // technique
				card.TrainingDexterity += int32(paramCell.Value)
			default:
				panic("Unexpected training content type")
			}
		case enum.TrainingTreeCellTypeVoice:
			naviActionId := trainingTree.NaviActionIds[cell.TrainingContentNo]
			session.UpdateVoice(naviActionId, true)
		case enum.TrainingTreeCellTypeStory:
			// training_content_type 11 in m_training_tree_card_story_side
			storySideId, exist := trainingTree.TrainingTreeCardStorySides[int(enum.TrainingContentTypeStory)]
			if !exist {
				panic("story doesn't exist")
			}
			session.InsertStorySide(storySideId)
		case enum.TrainingTreeCellTypeAwakening:
			// idolize
			card.IsAwakening = true
			card.IsAwakeningImage = true
			storySideId, exist := trainingTree.TrainingTreeCardStorySides[int(enum.TrainingContentTypeAwakening)]
			if exist {
				session.InsertStorySide(storySideId)
			}
		case enum.TrainingTreeCellTypeSuit:
			// alternative suit is awarded based on amount of tile instead
			user_suit.InsertUserSuit(session, trainingTree.SuitMIds[cell.TrainingContentNo])
		case enum.TrainingTreeCellTypeCardActiveSkillOriginIncrease: // skill
			card.ActiveSkillLevel++
		case enum.TrainingTreeCellTypeCardPassiveSkillAdditionalExpansionSlot: // insight
			card.MaxFreePassiveSkill++
		case enum.TrainingTreeCellTypeCardPassiveSkillOriginIncrease: // ability
			card.PassiveSkillALevel++
		}
	}

	// progress reward
	for _, reward := range trainingTree.TrainingTreeProgressRewards {
		if reward.ActivateNum > int(card.TrainingActivatedCellCount)+req.CellMasterIds.Size() {
			break
		}
		if reward.ActivateNum > int(card.TrainingActivatedCellCount) {
			user_content.AddContent(session, reward.Reward)
		}
	}

	card.TrainingActivatedCellCount += int32(req.CellMasterIds.Size())

	if card.TrainingActivatedCellCount+1 == int32(len(cellContents)) {
		card.IsAllTrainingActivated = true
		member := user_member.GetMember(session, *masterCard.MemberMasterId)
		user_member.UpdateMember(session, member)
	}

	session.UpdateUserCard(card)

	// set "user_card_training_tree_cell_list" to the cell unlocked and insert the cell to db
	unlockedCells := []client.UserCardTrainingTreeCell{}
	for _, cellId := range req.CellMasterIds.Slice {
		unlockedCells = append(unlockedCells,
			client.UserCardTrainingTreeCell{
				CellId:      cellId,
				ActivatedAt: session.Time.Unix()})
	}

	session.InsertTrainingTreeCells(req.CardMasterId, unlockedCells)
	session.Finalize()

	common.JsonResponse(ctx, &response.ActivateTrainingTreeCellResponse{
		UserCardTrainingTreeCellList: session.GetTrainingTree(req.CardMasterId),
		UserModelDiff:                &session.UserModel,
	})
}

func init() {
	// TODO(refactor): move to individual files. 
	router.AddHandler("/trainingTree/fetchTrainingTree", FetchTrainingTree)
	router.AddHandler("/trainingTree/levelUpCard", LevelUpCard)
	router.AddHandler("/trainingTree/gradeUpCard", GradeUpCard)
	router.AddHandler("/trainingTree/activateTrainingTreeCell", ActivateTrainingTreeCell)
}
