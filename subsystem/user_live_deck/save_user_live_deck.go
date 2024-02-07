package user_live_deck

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_live_party"
	"elichika/subsystem/user_suit"
	"elichika/userdata"

	"reflect"
)

func SaveUserLiveDeck(session *userdata.Session, deckId int32,
	cardWithSuit generic.Dictionary[int32, generic.Nullable[int32]],
	squadDict generic.Dictionary[int32, client.LiveSquad]) {

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseDeckEdit {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseSuitChange
	}

	userLiveDeck := GetUserLiveDeck(session, deckId)
	for position, cardMasterId := range cardWithSuit.OrderedKey {
		suitMasterId := *cardWithSuit.GetOnly(cardMasterId)
		if !suitMasterId.HasValue {
			suitMasterId.HasValue = true
			// TODO(suit): this assume that the default suit for a card, if it has one, use the same id
			// this is true for existing database but might not be
			suitMasterId.Value = cardMasterId
			if !user_suit.HasSuit(session, cardMasterId) {
				suitMasterId.Value = session.Gamedata.Card[cardMasterId].Member.MemberInit.SuitMasterId
			}
		}
		reflect.ValueOf(&userLiveDeck).Elem().Field(position + 2).Set(reflect.ValueOf(generic.NewNullable(cardMasterId)))
		reflect.ValueOf(&userLiveDeck).Elem().Field(position + 2 + 9).Set(reflect.ValueOf(suitMasterId))
	}
	UpdateUserLiveDeck(session, userLiveDeck)
	for partyId, liveSquad := range squadDict.Map {
		userLiveParty := client.UserLiveParty{
			PartyId:        partyId,
			UserLiveDeckId: deckId,
		}
		userLiveParty.IconMasterId, userLiveParty.Name.DotUnderText = session.Gamedata.GetLivePartyInfoByCardMasterIds(
			liveSquad.CardMasterIds.Slice[0], liveSquad.CardMasterIds.Slice[1], liveSquad.CardMasterIds.Slice[2])
		for position := 0; position < 3; position++ {
			reflect.ValueOf(&userLiveParty).Elem().Field(position + 4).Set(
				reflect.ValueOf(generic.NewNullable(liveSquad.CardMasterIds.Slice[position])))
			reflect.ValueOf(&userLiveParty).Elem().Field(position + 4 + 3).Set(
				reflect.ValueOf(liveSquad.UserAccessoryIds.Slice[position]))
		}
		user_live_party.UpdateUserLiveParty(session, userLiveParty)
	}

}
