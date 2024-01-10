package userdata

import (
	"elichika/client"
	"elichika/utils"
)

func (session *Session) GetUserLiveDeck(userLiveDeckId int) client.UserLiveDeck {
	liveDeck := client.UserLiveDeck{}
	exist, err := session.Db.Table("u_live_deck").
		Where("user_id = ? AND user_live_deck_id = ?", session.UserId, userLiveDeckId).
		Get(&liveDeck)
	if err != nil {
		panic(err)
	}
	if !exist {
		panic("Deck doesn't exist")
	}
	return liveDeck
}

func (session *Session) UpdateUserLiveDeck(liveDeck client.UserLiveDeck) {
	session.UserModel.UserLiveDeckById.Set(liveDeck.UserLiveDeckId, liveDeck)
}

func liveDeckFinalizer(session *Session) {
	for _, deck := range session.UserModel.UserLiveDeckById.Map {
		affected, err := session.Db.Table("u_live_deck").
			Where("user_id = ? AND user_live_deck_id = ?", session.UserId, deck.UserLiveDeckId).AllCols().
			Update(deck)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_live_deck", deck)
		}
	}
}

func (session *Session) InsertLiveDecks(decks []client.UserLiveDeck) {
	for _, deck := range decks {
		session.UpdateUserLiveDeck(deck)
	}
}

func init() {
	addFinalizer(liveDeckFinalizer)
	addGenericTableFieldPopulator("u_live_deck", "UserLiveDeckById")
}
