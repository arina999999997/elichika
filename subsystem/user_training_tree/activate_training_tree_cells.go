package user_training_tree

import (
	"elichika/enum"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_mission"
	"elichika/subsystem/user_story_side"
	"elichika/subsystem/user_suit"
	"elichika/subsystem/user_voice"
	"elichika/userdata"
)

// return the training tree for a card
func ActivateTrainingTreeCells(session *userdata.Session, cardMasterId int32, cellIds []int32) {

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseTrainingActivateCell {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseDeckEdit
	}

	card := user_card.GetUserCard(session, cardMasterId)
	masterCard := session.Gamedata.Card[cardMasterId]
	trainingTree := masterCard.TrainingTree
	cellContents := trainingTree.TrainingTreeMapping.TrainingTreeCellContents
	for _, cellId := range cellIds {
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
			user_voice.UpdateUserVoice(session, naviActionId, true)
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
			// mision
			user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountSchoolIdolAwakening, nil, nil,
				user_mission.AddProgressHandler, int32(1))
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
		if reward.ActivateNum > int(card.TrainingActivatedCellCount)+len(cellIds) {
			break
		}
		if reward.ActivateNum > int(card.TrainingActivatedCellCount) {
			user_content.AddContent(session, reward.Reward)
		}
	}

	card.TrainingActivatedCellCount += int32(len(cellIds))

	if card.TrainingActivatedCellCount+1 == int32(len(cellContents)) {
		card.IsAllTrainingActivated = true
	}

	user_card.UpdateUserCard(session, card)

	for _, cellId := range cellIds {
		type Wrapper struct {
			CardMasterId int32
			CellId       int32
			ActivatedAt  int64
		}
		userdata.GenericDatabaseInsert(session, "u_card_training_tree_cell", Wrapper{
			CardMasterId: cardMasterId,
			CellId:       cellId,
			ActivatedAt:  session.Time.Unix(),
		})
	}
}
