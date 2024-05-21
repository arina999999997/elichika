package client

type WeekdayState struct {
	Weekday       int32 `json:"weekday" enum:"Weekday"`
	NextWeekdayAt int64 `json:"next_weekday_at"`
}
