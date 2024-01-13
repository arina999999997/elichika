package client

type DailyTheaterDetail struct {
	DailyTheaterId int32         `json:"daily_theater_id"`
	Title          LocalizedText `json:"title"`
	DetailText     LocalizedText `json:"detail_text"`
	Year           int32         `json:"year"`
	Month          int32         `json:"month"`
	Day            int32         `json:"day"`
}
