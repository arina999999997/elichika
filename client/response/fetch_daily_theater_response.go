package response

import (
	"elichika/client"
)

type FetchDailyTheaterResponse struct {
	DailyTheaterDetail client.DailyTheaterDetail `json:"daily_theater_detail"`
	UserModelDiff      *client.UserModel         `json:"user_model_diff"`
}
