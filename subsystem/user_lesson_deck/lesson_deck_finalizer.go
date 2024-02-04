package user_lesson_deck

import (
	"elichika/userdata"
	"elichika/utils"
)

func lessonDeckFinalizer(session *userdata.Session) {
	for _, deck := range session.UserModel.UserLessonDeckById.Map {
		affected, err := session.Db.Table("u_lesson_deck").
			Where("user_id = ? AND user_lesson_deck_id = ?", session.UserId, deck.UserLessonDeckId).AllCols().
			Update(*deck)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_lesson_deck", *deck)
		}
	}
}

func init() {
	userdata.AddFinalizer(lessonDeckFinalizer)
}
