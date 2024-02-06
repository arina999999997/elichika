package user_live_deck

import (
	"elichika/client"
	"elichika/userdata"
)

func UpdateUserLiveDeck(session *userdata.Session, liveDeck client.UserLiveDeck) {
	session.UserModel.UserLiveDeckById.Set(liveDeck.UserLiveDeckId, liveDeck)
}
