package client

type DailyTheaterArchiveMasterRow struct {
	DailyTheaterId int32         `json:"daily_theater_id"`
	Year           int32         `json:"year"`
	Month          int32         `json:"month"`
	Day            int32         `json:"day"`
	Title          LocalizedText `json:"title"`
	PublishedAt    int64         `json:"published_at"`
}
