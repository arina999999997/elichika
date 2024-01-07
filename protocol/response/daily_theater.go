package response

import (
	"elichika/model"
)

type DailyTheaterDetail struct {
	DailyTheaterId int                 `json:"daily_theater_id"`
	Title          model.LocalizedText `json:"title"`
	DetailText     model.LocalizedText `json:"detail_text"`
	Year           int                 `json:"year"`
	Month          int                 `json:"month"`
	Day            int                 `json:"day"`
}
type FetchDailyTheaterResponse struct {
	DailyTheaterDetail DailyTheaterDetail `json:"daily_theater_detail"`
	// UserModelDiff
}

type DailyTheaterArchiveMasterRow struct {
	DailyTheaterId int   `json:"daily_theater_id"`
	Year           int   `json:"year"`
	Month          int   `json:"month"`
	Day            int   `json:"day"`
	PublishedAt    int64 `json:"published_at"`
}

type DailyTheaterArchiveMemberMasterRow struct {
	DailyTheaterId int `json:"daily_theater_id"`
	MemberMasterId int `json:"member_master_id"`
}

type FetchDailyTheaterArchiveResponse struct {
	DailyTheaterArchiveMasterRows       []DailyTheaterArchiveMasterRow       `json:"daily_theater_archive_master_rows"`
	DailyTheaterArchiveMemberMasterRows []DailyTheaterArchiveMemberMasterRow `json:"daily_theater_archive_member_master_rows"`
}
