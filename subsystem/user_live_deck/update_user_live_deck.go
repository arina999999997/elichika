package user_live_deck

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"

	"reflect"
)

func UpdateUserLiveDeck(session *userdata.Session, deckId int32,
	cardWithSuit generic.Dictionary[int32, generic.Nullable[int32]],
	squadDict generic.Dictionary[int32, client.LiveSquad]) {

	userLiveDeck := session.GetUserLiveDeck(deckId)
	for position, cardMasterId := range cardWithSuit.Order {
		suitMasterId := *cardWithSuit.GetOnly(cardMasterId)
		if !suitMasterId.HasValue {
			// TODO(suit): maybe we can assign the suit of the card instead
			suitMasterId = generic.NewNullable(session.Gamedata.Card[cardMasterId].Member.MemberInit.SuitMasterId)
		}
		reflect.ValueOf(&userLiveDeck).Elem().Field(position + 2).Set(reflect.ValueOf(generic.NewNullable(cardMasterId)))
		reflect.ValueOf(&userLiveDeck).Elem().Field(position + 2 + 9).Set(reflect.ValueOf(suitMasterId))
	}
	session.UpdateUserLiveDeck(userLiveDeck)
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
		session.UpdateUserLiveParty(userLiveParty)
	}

}
