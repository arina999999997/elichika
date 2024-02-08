package training_tree

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_member"
	"elichika/subsystem/user_story_side"
	"elichika/subsystem/user_suit"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func activateTrainingTreeCell(ctx *gin.Context) {
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

	card := user_card.GetUserCard(session, req.CardMasterId)
	masterCard := session.Gamedata.Card[req.CardMasterId]
	trainingTree := masterCard.TrainingTree
	cellContents := trainingTree.TrainingTreeMapping.TrainingTreeCellContents
	for _, cellId := range req.CellMasterIds.Slice {
		cell := &cellContents[cellId]
		// consume practice items only if this is not tutorial
		if session.UserStatus.TutorialPhase == enum.TutorialPhaseTutorialEnd {
			for _, resource := range cell.TrainingTreeCellItemSet.Resources {
				user_content.RemoveContent(session, resource)
			}
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
			user_story_side.InsertStorySide(session, storySideId)
		case enum.TrainingTreeCellTypeAwakening:
			// idolize
			card.IsAwakening = true
			card.IsAwakeningImage = true
			storySideId, exist := trainingTree.TrainingTreeCardStorySides[int(enum.TrainingContentTypeAwakening)]
			if exist {
				user_story_side.InsertStorySide(session, storySideId)
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

	user_card.UpdateUserCard(session, card)

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
	router.AddHandler("/trainingTree/activateTrainingTreeCell", activateTrainingTreeCell)
}
