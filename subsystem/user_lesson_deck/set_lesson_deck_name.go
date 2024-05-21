package user_lesson_deck

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/userdata"
)

func SetLessonDeckName(session *userdata.Session, deckId int32, name string) (*response.UserModelResponse, *response.RecoverableExceptionResponse) {
	if session.Gamedata.NgWord.HasMatch(name) {
		return nil, &response.RecoverableExceptionResponse{
			RecoverableExceptionType: enum.RecoverableExceptionTypeCommonNgWord,
		}
	}
	lessonDeck := GetUserLessonDeck(session, deckId)
	lessonDeck.Name = name
	UpdateLessonDeck(session, lessonDeck)
	return &response.UserModelResponse{
		UserModel: &session.UserModel,
	}, nil
}
