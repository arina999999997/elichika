package user_live_deck

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_live_party"
	"elichika/subsystem/user_mission"
	"elichika/subsystem/user_suit"
	"elichika/userdata"

	"reflect"
)

func ChangeLiveDeckCards(session *userdata.Session, deckId int32, cardMasterIds generic.Dictionary[int32, int32]) {
	deck := GetUserLiveDeck(session, deckId)
	replacingCard := map[int32]int32{}
	if cardMasterIds.Size() == 1 { // a new card
		position := cardMasterIds.OrderedKey[0]
		oldCardMasterId := reflect.ValueOf(deck).Field(1 + int(position)).Interface().(generic.Nullable[int32]).Value
		newCardMasterId := *cardMasterIds.GetOnly(position)
		replacingCard[oldCardMasterId] = newCardMasterId
		reflect.ValueOf(&deck).Elem().Field(1 + int(position)).Set(reflect.ValueOf(generic.NewNullable(newCardMasterId)))
		// TODO(suit): this assume that the default suit for a card, if it has one, use the same id
		// this is true for existing database but might not be
		suitMasterId := newCardMasterId
		_, exist := session.Gamedata.Suit[suitMasterId]
		if (!exist) || (!user_suit.HasSuit(session, suitMasterId)) {
			suitMasterId = session.Gamedata.Card[newCardMasterId].Member.MemberInit.SuitMasterId
		}
		reflect.ValueOf(&deck).Elem().Field(1 + int(position) + 9).Set(reflect.ValueOf(generic.NewNullable(suitMasterId)))
	} else { // swap card
		position0 := cardMasterIds.OrderedKey[0]
		position1 := cardMasterIds.OrderedKey[1]
		cardMasterId0 := reflect.ValueOf(deck).Field(1 + int(position0)).Interface().(generic.Nullable[int32]).Value
		cardMasterId1 := reflect.ValueOf(deck).Field(1 + int(position1)).Interface().(generic.Nullable[int32]).Value
		replacingCard[cardMasterId0] = cardMasterId1
		replacingCard[cardMasterId1] = cardMasterId0
		suitMasterId0 := reflect.ValueOf(deck).Field(1 + int(position0) + 9).Interface().(generic.Nullable[int32]).Value
		suitMasterId1 := reflect.ValueOf(deck).Field(1 + int(position1) + 9).Interface().(generic.Nullable[int32]).Value
		if (cardMasterId0 != *cardMasterIds.GetOnly(position1)) || (cardMasterId1 != *cardMasterIds.GetOnly(position0)) {
			panic("unexpected card swap")
		}
		reflect.ValueOf(&deck).Elem().Field(1 + int(position0)).Set(reflect.ValueOf(generic.NewNullable(cardMasterId1)))
		reflect.ValueOf(&deck).Elem().Field(1 + int(position0) + 9).Set(reflect.ValueOf(generic.NewNullable(suitMasterId1)))
		reflect.ValueOf(&deck).Elem().Field(1 + int(position1)).Set(reflect.ValueOf(generic.NewNullable(cardMasterId0)))
		reflect.ValueOf(&deck).Elem().Field(1 + int(position1) + 9).Set(reflect.ValueOf(generic.NewNullable(suitMasterId0)))
	}
	UpdateUserLiveDeck(session, deck)

	parties := map[int32]client.UserLiveParty{}
	for replaced := range replacingCard {
		party := user_live_party.GetUserLivePartyWithDeckAndCardId(session, deckId, replaced)
		parties[party.PartyId] = party
	}

	for _, party := range parties {
		for i := 1; i <= 3; i++ {
			currentCardMasterId := reflect.ValueOf(party).Field(3 + i).Interface().(generic.Nullable[int32]).Value
			newCardMasterId, exist := replacingCard[currentCardMasterId]
			if exist {
				reflect.ValueOf(&party).Elem().Field(3 + i).Set(reflect.ValueOf(generic.NewNullable(newCardMasterId)))
			}
		}

		party.IconMasterId, party.Name.DotUnderText = session.Gamedata.GetLivePartyInfoByCardMasterIds(
			party.CardMasterId1.Value, party.CardMasterId2.Value, party.CardMasterId3.Value)
		user_live_party.UpdateUserLiveParty(session, party)
	}

	// mission progress tracking
	user_mission.UpdateProgress(session, enum.MissionClearConditionTypeCountEditLiveDeck, nil, nil,
		user_mission.AddProgressHandler, int32(1))
}
