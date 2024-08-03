package event_marathon_ranking

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_event/marathon"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// response: FetchEventMarathonRankingResponse
// alternative response: RecoverableExceptionResponse
func fetchEventMarathonRanking(ctx *gin.Context) {
	req := request.FetchEventMarathonRankingRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	success, failure := marathon.FetchEventMarathonRanking(session, req.EventId)
	if success != nil {
		common.JsonResponse(ctx, success)
	} else {
		common.AlternativeJsonResponse(ctx, failure)
	}
}

func init() {
	router.AddHandler("/", "POST", "/eventMarathonRanking/fetchEventMarathonRanking", fetchEventMarathonRanking)
}
