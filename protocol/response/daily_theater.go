package response

type DailyTheaterDetail struct {
	DailyTheaterID int           `json:"daily_theater_id"`
	Title          LocalizedText `json:"title"`
	DetailText     LocalizedText `json:"detail_text"`
	Year           int           `json:"year"`
	Month          int           `json:"month"`
	Day            int           `json:"day"`
}
type FetchDailyTheaterResponse struct {
	DailyTheaterDetail DailyTheaterDetail `json:"daily_theater_detail"`
	// UserModelDiff
}

type DailyTheaterArchiveMasterRow struct {
	DailyTheaterID int   `json:"daily_theater_id"`
	Year           int   `json:"year"`
	Month          int   `json:"month"`
	Day            int   `json:"day"`
	PublishedAt    int64 `json:"published_at"`
}

type DailyTheaterArchiveMemberMasterRow struct {
	DailyTheaterID int `json:"daily_theater_id"`
	MemberMasterID int `json:"member_master_id"`
}

type FetchDailyTheaterArchiveResponse struct {
	DailyTheaterArchiveMasterRows       []DailyTheaterArchiveMasterRow       `json:"daily_theater_archive_master_rows"`
	DailyTheaterArchiveMemberMasterRows []DailyTheaterArchiveMemberMasterRow `json:"daily_theater_archive_member_master_rows"`
}
