package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetUserLiveDeck(userLiveDeckId int) model.UserLiveDeck {
	liveDeck := model.UserLiveDeck{}
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

func (session *Session) UpdateUserLiveDeck(liveDeck model.UserLiveDeck) {
	session.UserLiveDeckMapping.SetList(&session.UserModel.UserLiveDeckById).Update(liveDeck)
}

func liveDeckFinalizer(session *Session) {
	for _, deck := range session.UserModel.UserLiveDeckById.Objects {
		affected, err := session.Db.Table("u_live_deck").
			Where("user_id = ? AND user_live_deck_id = ?", session.UserId, deck.UserLiveDeckId).AllCols().
			Update(deck)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_live_deck", deck)
		}
	}
}

func (session *Session) InsertLiveDecks(decks []model.UserLiveDeck) {
	session.UserModel.UserLiveDeckById.Objects = append(session.UserModel.UserLiveDeckById.Objects, decks...)
}

func init() {
	addFinalizer(liveDeckFinalizer)
	addGenericTableFieldPopulator("u_live_deck", "UserLiveDeckById")
}
