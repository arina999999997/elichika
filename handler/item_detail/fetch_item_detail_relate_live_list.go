package item_detail

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/time"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

// TODO(extra): Implement live campaign
func fetchItemDetailRelateLiveList(ctx *gin.Context) {
	// there is no request body

	session := ctx.MustGet("session").(*userdata.Session)
	common.JsonResponse(ctx, response.FetchItemDetailRelateLiveListResponse{
		WeekdayState: time.GetWeekdayState(session),
	})
}

func init() {
	router.AddHandler("/", "POST", "/itemDetail/fetchItemDetailRelateLiveList", fetchItemDetailRelateLiveList)
}
