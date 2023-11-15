package userdata

import (
	"elichika/model"

	"fmt"
)

func (session *Session) GetLessonDeck(userLessonDeckId int) model.UserLessonDeck {
	deck, exist := session.UserLessonDeckDiffs[userLessonDeckId]
	if exist {
		return deck
	}
	deck = model.UserLessonDeck{}
	exists, err := session.Db.Table("u_lesson_deck").
		Where("user_id = ? AND user_lesson_deck_id = ?", session.UserStatus.UserID, userLessonDeckId).
		Get(&deck)
	if err != nil {
		panic(err)
	}
	if !exists {
		panic("deck not found")
	}
	return deck
}

func (session *Session) UpdateLessonDeck(userLessonDeck model.UserLessonDeck) {
	session.UserLessonDeckDiffs[userLessonDeck.UserLessonDeckID] = userLessonDeck
}

func (session *Session) FinalizeUserLessonDeckDiffs() []any {
	userLessonDeckByID := []any{}
	for userLessonDeckID, userLessonDeck := range session.UserLessonDeckDiffs {
		session.UserModelCommon.UserLessonDeckByID.PushBack(userLessonDeck)
		userLessonDeckByID = append(userLessonDeckByID, userLessonDeckID)
		userLessonDeckByID = append(userLessonDeckByID, userLessonDeck)
		affected, err := session.Db.Table("u_lesson_deck").
			Where("user_id = ? AND user_lesson_deck_id = ?", session.UserStatus.UserID, userLessonDeckID).
			AllCols().Update(userLessonDeck)
		if (err != nil) || (affected != 1) {
			panic(err)
		}
	}
	return userLessonDeckByID
}

func (session *Session) GetAllLessonDecks() []model.UserLessonDeck {
	decks := []model.UserLessonDeck{}
	err := session.Db.Table("u_lesson_deck").Where("user_id = ?", session.UserStatus.UserID).Find(&decks)
	if err != nil {
		panic(err)
	}
	return decks
}

func (session *Session) InsertLessonDecks(decks []model.UserLessonDeck) {
	count, err := session.Db.Table("u_lesson_deck").Insert(&decks)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted ", count, " lesson decks")
}
