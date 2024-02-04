package user_lesson_deck

import (
	"elichika/userdata"
)

func SetLessonDeckName(session *userdata.Session, deckId int32, name string) {
	lessonDeck := GetUserLessonDeck(session, deckId)
	lessonDeck.Name = name
	UpdateLessonDeck(session, lessonDeck)
}
