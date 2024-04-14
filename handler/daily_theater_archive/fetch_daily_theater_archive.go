package daily_theater_archive

import (
	"elichika/client"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"

	"github.com/gin-gonic/gin"
)

func fetchDailyTheaterArchive(ctx *gin.Context) {
	// this is used to publish new daily theater without having to update the database
	// client have the old items in m_daily_theater_archive_client and m_daily_theater_archive_member_client
	// client's missing 20230629 and 20230630

	// There is no request body

	resp := response.FetchDailyTheaterArchiveResponse{}
	// this isn't the actual thing
	resp.DailyTheaterArchiveMasterRows.Append(
		client.DailyTheaterArchiveMasterRow{
			DailyTheaterId: 1001243,
			Year:           2023,
			Month:          6,
			Day:            29,
			PublishedAt:    1687964400,
		})
	resp.DailyTheaterArchiveMemberMasterRows.Append(client.DailyTheaterArchiveMemberMasterRow{
		DailyTheaterId: 1001243,
		MemberMasterId: 101, // Chika
	})
	resp.DailyTheaterArchiveMemberMasterRows.Append(client.DailyTheaterArchiveMemberMasterRow{
		DailyTheaterId: 1001243,
		MemberMasterId: 106, // Yoshiko
	})

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/dailyTheaterArchive/fetchDailyTheaterArchive", fetchDailyTheaterArchive)
}
