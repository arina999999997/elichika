package client

type WeekdayState struct {
	Weekday       int32 `json:"weekday"`
	NextWeekdayAt int64 `json:"next_weekday_at"`
}
