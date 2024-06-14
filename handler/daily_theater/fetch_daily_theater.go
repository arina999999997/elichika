package daily_theater

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_daily_theater"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func fetchDailyTheater(ctx *gin.Context) {
	req := request.FetchDailyTheaterRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	successResponse, failureRespose := user_daily_theater.FetchDailyTheater(session, req.DailyTheaterId)
	if successResponse == nil {
		common.AlternativeJsonResponse(ctx, failureRespose)
	} else {
		common.JsonResponse(ctx, successResponse)
	}
}

func init() {
	router.AddHandler("/", "POST", "/dailyTheater/fetchDailyTheater", fetchDailyTheater)
}
