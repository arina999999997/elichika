package user_daily_theater

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/enum"
	"elichika/generic"
	"elichika/userdata"
)

func FetchDailyTheater(session *userdata.Session, id generic.Nullable[int32]) (*response.FetchDailyTheaterResponse, *response.RecoverableExceptionResponse) {
	if !id.HasValue {
		id.Value = session.Gamedata.LastestDailyTheaterId
	}
	dailyTheater, exist := session.Gamedata.DailyTheater[id.Value]
	if exist {
		return &response.FetchDailyTheaterResponse{
			DailyTheaterDetail: client.DailyTheaterDetail{
				DailyTheaterId: dailyTheater.DailyTheaterId,
				Title:          dailyTheater.Title,
				DetailText:     dailyTheater.DetailText,
				Year:           dailyTheater.Year,
				Month:          dailyTheater.Month,
				Day:            dailyTheater.Day,
			},
			UserModelDiff: &session.UserModel,
		}, nil
	} else {
		return nil, &response.RecoverableExceptionResponse{
			RecoverableExceptionType: enum.RecoverableExceptionTypeDailyTheaterOutOfTerm,
		}
	}
}
