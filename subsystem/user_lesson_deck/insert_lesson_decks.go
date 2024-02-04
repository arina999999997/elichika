package user_lesson_deck

import (
	"elichika/client"
	"elichika/userdata"
)

func InsertLessonDecks(session *userdata.Session, decks []client.UserLessonDeck) {
	for _, deck := range decks {
		UpdateLessonDeck(session, deck)
	}
}
