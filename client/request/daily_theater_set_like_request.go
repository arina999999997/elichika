package request

type DailyTheaterSetLikeRequest struct {
	DailyTheaterId int32 `json:"daily_theater_id"`
	IsLike         bool  `json:"is_like"`
}
