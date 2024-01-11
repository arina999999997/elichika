package userdata

import (
	"elichika/client"
	"elichika/utils"
)

func (session *Session) GetUserLessonDeck(userLessonDeckId int32) client.UserLessonDeck {
	ptr, exist := session.UserModel.UserLessonDeckById.Get(userLessonDeckId)
	if exist {
		return *ptr
	}
	deck := client.UserLessonDeck{}
	exist, err := session.Db.Table("u_lesson_deck").
		Where("user_id = ? AND user_lesson_deck_id = ?", session.UserId, userLessonDeckId).
		Get(&deck)
	utils.CheckErr(err)
	if !exist {
		panic("deck not found")
	}
	return deck
}

func (session *Session) UpdateLessonDeck(userLessonDeck client.UserLessonDeck) {
	session.UserModel.UserLessonDeckById.Set(userLessonDeck.UserLessonDeckId, userLessonDeck)
}

func lessonDeckFinalizer(session *Session) {
	for _, deck := range session.UserModel.UserLessonDeckById.Map {
		affected, err := session.Db.Table("u_lesson_deck").
			Where("user_id = ? AND user_lesson_deck_id = ?", session.UserId, deck.UserLessonDeckId).AllCols().
			Update(*deck)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_lesson_deck", *deck)
		}
	}
}

func (session *Session) InsertLessonDecks(decks []client.UserLessonDeck) {
	for _, deck := range decks {
		session.UpdateLessonDeck(deck)
	}
}

func init() {
	addFinalizer(lessonDeckFinalizer)
	addGenericTableFieldPopulator("u_lesson_deck", "UserLessonDeckById")
}
