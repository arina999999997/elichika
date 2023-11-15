package userdata

import (
	"elichika/model"

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
		session.UserModelCommon.UserLiveDeckByID.PushBack(userLiveDeck)
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
	if err != nil {
		panic(err)
	}
	return decks
}

func (session *Session) InsertLiveDecks(decks []model.UserLiveDeck) {
	count, err := session.Db.Table("u_live_deck").Insert(&decks)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted ", count, " live decks")
}
