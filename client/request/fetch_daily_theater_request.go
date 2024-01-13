package request

import (
	"elichika/generic"
)

type FetchDailyTheaterRequest struct {
	DailyTheaterId generic.Nullable[int32] `json:"daily_theater_id"`
}
