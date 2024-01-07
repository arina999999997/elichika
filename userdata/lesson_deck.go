package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetUserLessonDeck(userLessonDeckId int) model.UserLessonDeck {
	pos, exist := session.UserLessonDeckMapping.SetList(&session.UserModel.UserLessonDeckById).Map[int64(userLessonDeckId)]
	if exist {
		return session.UserModel.UserLessonDeckById.Objects[pos]
	}
	deck := model.UserLessonDeck{}
	exist, err := session.Db.Table("u_lesson_deck").
		Where("user_id = ? AND user_lesson_deck_id = ?", session.UserId, userLessonDeckId).
		Get(&deck)
	utils.CheckErr(err)
	if !exist {
		panic("deck not found")
	}
	return deck
}

func (session *Session) UpdateLessonDeck(userLessonDeck model.UserLessonDeck) {
	session.UserLessonDeckMapping.SetList(&session.UserModel.UserLessonDeckById).Update(userLessonDeck)
}

func lessonDeckFinalizer(session *Session) {
	for _, deck := range session.UserModel.UserLessonDeckById.Objects {
		affected, err := session.Db.Table("u_lesson_deck").
			Where("user_id = ? AND user_lesson_deck_id = ?", session.UserId, deck.UserLessonDeckId).AllCols().
			Update(deck)
		utils.CheckErr(err)
		if affected == 0 {
			genericDatabaseInsert(session, "u_lesson_deck", deck)
		}
	}
}

func (session *Session) InsertLessonDecks(decks []model.UserLessonDeck) {
	session.UserModel.UserLessonDeckById.Objects = append(session.UserModel.UserLessonDeckById.Objects, decks...)
}

func init() {
	addFinalizer(lessonDeckFinalizer)
	addGenericTableFieldPopulator("u_lesson_deck", "UserLessonDeckById")
}
