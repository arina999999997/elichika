package user_daily_theater

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/userdata"
)

// TODO(extra): implement daily theater adding + reload in admin webui
// and maybe preloaded theater to be released later
var dailyTheaterArchive = map[string]*response.FetchDailyTheaterArchiveResponse{}

func FetchDailyTheaterArchive(session *userdata.Session) *response.FetchDailyTheaterArchiveResponse {
	archive, exist := dailyTheaterArchive[session.Gamedata.Language]
	if exist {
		return archive
	}
	archive = &response.FetchDailyTheaterArchiveResponse{}
	for _, dailyTheater := range session.Gamedata.DailyTheater {
		if dailyTheater.IsInClient {
			continue
		}
		archive.DailyTheaterArchiveMasterRows.Append(client.DailyTheaterArchiveMasterRow{
			DailyTheaterId: dailyTheater.DailyTheaterId,
			Year:           dailyTheater.Year,
			Month:          dailyTheater.Month,
			Day:            dailyTheater.Day,
			Title:          dailyTheater.Title,
			PublishedAt:    dailyTheater.PublishedAt,
		})
		for _, memberId := range dailyTheater.Members {
			archive.DailyTheaterArchiveMemberMasterRows.Append(client.DailyTheaterArchiveMemberMasterRow{
				DailyTheaterId: dailyTheater.DailyTheaterId,
				MemberMasterId: memberId,
			})
		}
	}
	dailyTheaterArchive[session.Gamedata.Language] = archive
	return archive
}
