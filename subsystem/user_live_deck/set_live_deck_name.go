package user_live_deck

import (
	"elichika/userdata"
)

func SetLiveDeckName(session *userdata.Session, deckId int32, deckName string) {
	liveDeck := GetUserLiveDeck(session, deckId)
	liveDeck.Name.DotUnderText = deckName
	UpdateUserLiveDeck(session, liveDeck)
}
