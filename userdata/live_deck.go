package userdata

import (
	"elichika/model"
	"elichika/utils"

	"fmt"
)

func (session *Session) GetUserLiveDeck(userLiveDeckID int) model.UserLiveDeck {
	liveDeck := model.UserLiveDeck{}
	exists, err := session.Db.Table("u_live_deck").
		Where("user_id = ? AND user_live_deck_id = ?", session.UserStatus.UserID, userLiveDeckID).
		Get(&liveDeck)
	if err != nil {
		panic(err)
	}
	if !exists {
		panic("Deck doesn't exist")
	}
	return liveDeck
}

func (session *Session) UpdateUserLiveDeck(liveDeck model.UserLiveDeck) {
	session.UserLiveDeckDiffs[liveDeck.UserLiveDeckID] = liveDeck
}

func (session *Session) FinalizeUserLiveDeckDiffs() []any {
	userLiveDeckByID := []any{}
	for userLiveDeckId, userLiveDeck := range session.UserLiveDeckDiffs {
		session.UserModel.UserLiveDeckByID.PushBack(userLiveDeck)
		userLiveDeckByID = append(userLiveDeckByID, userLiveDeckId)
		userLiveDeckByID = append(userLiveDeckByID, userLiveDeck)
		affected, err := session.Db.Table("u_live_deck").
			Where("user_id = ? AND user_live_deck_id = ?", session.UserStatus.UserID, userLiveDeckId).
			AllCols().Update(userLiveDeck)
		if (err != nil) || (affected != 1) {
			panic(err)
		}
	}
	return userLiveDeckByID
}

func (session *Session) GetAllLiveDecks() []model.UserLiveDeck {
	decks := []model.UserLiveDeck{}
	err := session.Db.Table("u_live_deck").Where("user_id = ?", session.UserStatus.UserID).Find(&decks)
	utils.CheckErr(err)
	return decks
}

func (session *Session) InsertLiveDecks(decks []model.UserLiveDeck) {
	count, err := session.Db.Table("u_live_deck").Insert(&decks)
	utils.CheckErr(err)
	fmt.Println("Inserted ", count, " live decks")
}

func init() {
	addGenericTableFieldPopulator("u_live_deck", "UserLiveDeckByID")
}
