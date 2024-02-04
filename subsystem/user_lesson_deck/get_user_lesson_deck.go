package user_lesson_deck

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetUserLessonDeck(session *userdata.Session, userLessonDeckId int32) client.UserLessonDeck {
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
