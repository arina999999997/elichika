package daily_theater_archive

import (
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_daily_theater"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchDailyTheaterArchive(ctx *gin.Context) {
	// this is used to publish new daily theater without having to update the database
	// client have the old items in m_daily_theater_archive_client and m_daily_theater_archive_member_client
	// as of EOS, the client is missing 20230629 and 20230630

	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)
	common.JsonResponse(ctx, user_daily_theater.FetchDailyTheaterArchive(session))
}

func init() {
	router.AddHandler("/", "POST", "/dailyTheaterArchive/fetchDailyTheaterArchive", fetchDailyTheaterArchive)
}
