package live

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/time"
	"elichika/userdata"
	"elichika/utils"

	"github.com/gin-gonic/gin"
)

func resume(ctx *gin.Context) {
	// there is no request body

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	exist, live, startReq := session.LoadUserLive()
	utils.MustExist(exist)

	common.JsonResponse(ctx, &response.ResumeLiveResponse{
		Live:          live,
		PartnerUserId: startReq.PartnerUserId,
		IsAutoPlay:    startReq.IsAutoPlay,
		WeekdayState:  time.GetWeekdayState(session),
	})
}

func init() {
	router.AddHandler("/live/resume", resume)
}
