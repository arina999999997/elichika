package user_training_tree

import (
	"elichika/enum"
	"elichika/item"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_content"
	"elichika/userdata"
)

func LevelUpCard(session *userdata.Session, cardMasterId, additionalLevel int32) {
	if session.UserStatus.TutorialPhase == enum.TutorialPhaseTrainingLevelUp {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseTrainingActivateCell
	}

	cardLevel := session.Gamedata.CardLevel[session.Gamedata.Card[cardMasterId].CardRarityType]
	card := user_card.GetUserCard(session, cardMasterId)
	user_content.RemoveContent(session, item.Gold.Amount(int32(
		cardLevel.GameMoneyPrefixSum[card.Level+additionalLevel]-cardLevel.GameMoneyPrefixSum[card.Level])))
	user_content.RemoveContent(session, item.EXP.Amount(int32(
		cardLevel.ExpPrefixSum[card.Level+additionalLevel]-cardLevel.ExpPrefixSum[card.Level])))
	card.Level += additionalLevel
	user_card.UpdateUserCard(session, card)
}
