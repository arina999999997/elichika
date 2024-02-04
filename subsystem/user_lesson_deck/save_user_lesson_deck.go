package user_lesson_deck

import (
	"elichika/client/request"
	"elichika/userdata"

	"reflect"
)

func SaveUserLessonDeck(session *userdata.Session, req request.SaveLessonDeckRequest) {
	userLessonDeck := GetUserLessonDeck(session, req.DeckId)
	for position, cardMasterId := range req.CardMasterIds.Map {
		reflect.ValueOf(&userLessonDeck).Elem().Field(int(position) + 1).Set(reflect.ValueOf(*cardMasterId))
	}
	UpdateLessonDeck(session, userLessonDeck)
}
