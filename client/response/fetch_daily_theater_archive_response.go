package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchDailyTheaterArchiveResponse struct {
	DailyTheaterArchiveMasterRows       generic.List[client.DailyTheaterArchiveMasterRow]       `json:"daily_theater_archive_master_rows"`
	DailyTheaterArchiveMemberMasterRows generic.List[client.DailyTheaterArchiveMemberMasterRow] `json:"daily_theater_archive_member_master_rows"`
}
