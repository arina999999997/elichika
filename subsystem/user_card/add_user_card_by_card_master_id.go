package user_card

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/item"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_info_trigger"
	"elichika/subsystem/user_member"
	"elichika/userdata"
)

// this function add a copy of a card to the user, it return a Nullable[Content] of the extra reward
// if that card is already maxed.
// Note that this assume the card is added, so it's not used for present box and for gacha retry
// the maxed limit break reward is also added directly to the user, the return value is only to help client display them
func AddUserCardByCardMasterId(session *userdata.Session, cardMasterId int32) client.AddedCardResult {
	card := GetUserCard(session, cardMasterId)
	masterCard := session.Gamedata.Card[cardMasterId]
	card.Grade++
	if card.Grade > enum.CardMaxGrade {
		// max limit already, award the item
		reward := item.SchoolIdolRadiance
		switch masterCard.CardRarityType {
		case enum.CardRarityTypeSRare:
			reward = reward.Amount(5)
		case enum.CardRarityTypeURare:
			reward = reward.Amount(25)
		}
		user_content.AddContent(session, reward)
		// not sure if this is the correct official server behavior or not, but it seems to work correctly
		return client.AddedCardResult{
			CardMasterId: cardMasterId,
			Level:        1,
			BeforeGrade:  enum.CardMaxGrade,
			AfterGrade:   enum.CardMaxGrade,
			// BeforeLoveLevelLimit: 0,
			// AfterLoveLevelLimit:  0,
			Content: generic.NewNullable(reward),
		}
	} else {
		beforeLoveLevelLimit, afterLoveLevelLimit := user_member.IncreaseMemberLoveLevelLimit(
			session, masterCard.Member.Id, masterCard.CardRarityType/10)
		beforeGrade := int32(0)
		if card.Grade > 0 { // is a limit break
			// add trigger card grade up so animation play when opening the card
			user_info_trigger.AddTriggerCardGradeUp(session, client.UserInfoTriggerCardGradeUp{
				CardMasterId:         card.CardMasterId,
				BeforeLoveLevelLimit: afterLoveLevelLimit, // this is correct
				AfterLoveLevelLimit:  afterLoveLevelLimit,
			})
			beforeGrade = card.Grade - 1
		}
		UpdateUserCard(session, card)
		return client.AddedCardResult{
			CardMasterId:         cardMasterId,
			Level:                1,
			BeforeGrade:          beforeGrade,
			AfterGrade:           card.Grade,
			BeforeLoveLevelLimit: beforeLoveLevelLimit,
			AfterLoveLevelLimit:  afterLoveLevelLimit,
		}
	}
}
