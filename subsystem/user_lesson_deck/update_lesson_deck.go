package user_lesson_deck

import (
	"elichika/client"
	"elichika/userdata"
)

func UpdateLessonDeck(session *userdata.Session, userLessonDeck client.UserLessonDeck) {
	session.UserModel.UserLessonDeckById.Set(userLessonDeck.UserLessonDeckId, userLessonDeck)
}
