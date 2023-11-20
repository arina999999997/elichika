package userdata

import (
	"elichika/model"
	"elichika/utils"

	"fmt"
)

func (session *Session) GetUserLessonDeck(userLessonDeckID int) model.UserLessonDeck {
	pos, exist := session.UserLessonDeckMapping.Map[int64(userLessonDeckID)]
	if exist {
		return session.UserModel.UserLessonDeckByID.Objects[pos]
	}
	deck := model.UserLessonDeck{}
	exists, err := session.Db.Table("u_lesson_deck").
		Where("user_id = ? AND user_lesson_deck_id = ?", session.UserStatus.UserID, userLessonDeckID).
		Get(&deck)
	utils.CheckErr(err)
	if !exists {
		panic("deck not found")
	}
	return deck
}

func (session *Session) UpdateLessonDeck(userLessonDeck model.UserLessonDeck) {
	session.UserLessonDeckMapping.SetList(&session.UserModel.UserLessonDeckByID).Update(userLessonDeck)
}

func lessonDeckFinalizer(session *Session) {
	for _, deck := range session.UserModel.UserLessonDeckByID.Objects {
		affected, err := session.Db.Table("u_lesson_deck").
			Where("user_id = ? AND user_lesson_deck_id = ?", session.UserStatus.UserID, deck.UserLessonDeckID).AllCols().
			Update(deck)
		utils.CheckErr(err)
		if affected == 0 {
			// all lesson deck must be inserted at login
			panic("user lesson deck doesn't exists")
		}
	}
}

func (session *Session) InsertLessonDecks(decks []model.UserLessonDeck) {
	count, err := session.Db.Table("u_lesson_deck").Insert(&decks)
	utils.CheckErr(err)
	fmt.Println("Inserted ", count, " lesson decks")
}

func init() {
	addFinalizer(lessonDeckFinalizer)
	addGenericTableFieldPopulator("u_lesson_deck", "UserLessonDeckByID")
}
