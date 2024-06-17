package user_training_tree

import (
	"elichika/config"
	"elichika/enum"
	"elichika/item"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
)

func LevelUpCard(session *userdata.Session, cardMasterId, additionalLevel int32) {
	if session.UserStatus.TutorialPhase == enum.TutorialPhaseTrainingLevelUp {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseTrainingActivateCell
	}
	masterCard := session.Gamedata.Card[cardMasterId]
	cardLevel := session.Gamedata.CardLevel[masterCard.CardRarityType]
	card := user_card.GetUserCard(session, cardMasterId)

	if config.Conf.ResourceConfig().ConsumePracticeItems {
		user_content.RemoveContent(session, item.Gold.Amount(int32(
			cardLevel.GameMoneyPrefixSum[card.Level+additionalLevel]-cardLevel.GameMoneyPrefixSum[card.Level])))
		user_content.RemoveContent(session, item.EXP.Amount(int32(
			cardLevel.ExpPrefixSum[card.Level+additionalLevel]-cardLevel.ExpPrefixSum[card.Level])))
	}
	card.Level += additionalLevel
	user_card.UpdateUserCard(session, card)

	// mission code
	if (card.Level >= masterCard.Rarity.MaxLevel) && (card.Level-additionalLevel < masterCard.Rarity.MaxLevel) {
		user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountSpecificRarityMakeLevelMax,
			&masterCard.CardRarityType, nil, user_mission.AddProgressHandler, int32(1))
	}

}
