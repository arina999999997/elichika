package user_training_tree

import (
	"elichika/client"
	"elichika/subsystem/user_card"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_info_trigger"
	"elichika/subsystem/user_member"
	"elichika/userdata"
)

func GradeUpCard(session *userdata.Session, cardMasterId, contentId int32) {
	masterCard := session.Gamedata.Card[cardMasterId]

	beforeLoveLevelLimit, afterLoveLevelLimit := user_member.IncreaseMemberLoveLevelLimit(
		session, masterCard.Member.Id, masterCard.Rarity.PlusLevel)

	card := user_card.GetUserCard(session, cardMasterId)
	card.Grade++
	user_card.UpdateUserCard(session, card)

	user_content.RemoveContent(session, masterCard.CardGradeUpItem[card.Grade][contentId])
	// we need to set user_info_trigger_card_grade_up_by_trigger_id
	// for the pop up after limit breaking
	// this trigger show the pop up after limit break
	user_info_trigger.AddTriggerCardGradeUp(session, client.UserInfoTriggerCardGradeUp{
		CardMasterId:         cardMasterId,
		BeforeLoveLevelLimit: beforeLoveLevelLimit,
		AfterLoveLevelLimit:  afterLoveLevelLimit,
	})
}
