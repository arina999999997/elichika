package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetUserLiveDeck(userLiveDeckID int) model.UserLiveDeck {
	liveDeck := model.UserLiveDeck{}
	exist, err := session.Db.Table("u_live_deck").
		Where("user_id = ? AND user_live_deck_id = ?", session.UserStatus.UserID, userLiveDeckID).
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
	session.UserLiveDeckMapping.SetList(&session.UserModel.UserLiveDeckByID).Update(liveDeck)
}

func liveDeckFinalizer(session *Session) {
	for _, deck := range session.UserModel.UserLiveDeckByID.Objects {
		affected, err := session.Db.Table("u_live_deck").
			Where("user_id = ? AND user_live_deck_id = ?", session.UserStatus.UserID, deck.UserLiveDeckID).AllCols().
			Update(deck)
		utils.CheckErr(err)
		if affected == 0 {
			_, err := session.Db.Table("u_live_deck").Insert(deck)
			utils.CheckErr(err)
		}
	}
}

func (session *Session) InsertLiveDecks(decks []model.UserLiveDeck) {
	session.UserModel.UserLiveDeckByID.Objects = append(session.UserModel.UserLiveDeckByID.Objects, decks...)
}

func init() {
	addFinalizer(liveDeckFinalizer)
	addGenericTableFieldPopulator("u_live_deck", "UserLiveDeckByID")
}
