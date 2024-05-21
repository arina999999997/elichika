package user_live_deck

import (
	"elichika/client/response"
	"elichika/enum"
	"elichika/userdata"
)

func SetLiveDeckName(session *userdata.Session, deckId int32, deckName string) (*response.UserModelResponse, *response.RecoverableExceptionResponse) {
	if session.Gamedata.NgWord.HasMatch(deckName) {
		return nil, &response.RecoverableExceptionResponse{
			RecoverableExceptionType: enum.RecoverableExceptionTypeCommonNgWord,
		}
	}
	liveDeck := GetUserLiveDeck(session, deckId)
	liveDeck.Name.DotUnderText = deckName
	UpdateUserLiveDeck(session, liveDeck)
	return &response.UserModelResponse{
		UserModel: &session.UserModel,
	}, nil
}
