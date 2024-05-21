package user_profile

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/userdata"
)

func SetProfile(session *userdata.Session, req request.SetUserProfileRequest) (*response.UserModelResponse, *response.RecoverableExceptionResponse) {
	if req.Name.HasValue {
		if session.Gamedata.NgWord.HasMatch(req.Name.Value) {
			return nil, &response.RecoverableExceptionResponse{
				RecoverableExceptionType: enum.RecoverableExceptionTypeNameContainsNgWord,
			}
		}
		session.UserStatus.Name.DotUnderText = req.Name.Value
	}
	if req.Nickname.HasValue {
		if session.Gamedata.NgWord.HasMatch(req.Nickname.Value) {
			return nil, &response.RecoverableExceptionResponse{
				RecoverableExceptionType: enum.RecoverableExceptionTypeNicknameContainsNgWord,
			}
		}
		session.UserStatus.Nickname.DotUnderText = req.Nickname.Value
	}
	if req.Message.HasValue {
		if session.Gamedata.NgWord.HasMatch(req.Message.Value) {
			return nil, &response.RecoverableExceptionResponse{
				RecoverableExceptionType: enum.RecoverableExceptionTypeMessageContainsNgWord,
			}
		}
		session.UserStatus.Message.DotUnderText = req.Message.Value
	}
	if req.DeviceToken.HasValue {
		session.UserStatus.DeviceToken = req.DeviceToken.Value
	}

	return &response.UserModelResponse{
		UserModel: &session.UserModel,
	}, nil
}
