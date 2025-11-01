package event_mining_ranking

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_event/mining"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// response: FetchEventMiningRankingResponse
// alternative response: RecoverableExceptionResponse
func fetchEventMiningRanking(ctx *gin.Context) {
	req := request.FetchEventMiningRankingRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	success, failure := mining.FetchEventMiningRanking(session, req.EventId)
	if success != nil {
		common.JsonResponse(ctx, success)
	} else {
		common.AlternativeJsonResponse(ctx, failure)
	}
}

func init() {
	router.AddHandler("/", "POST", "/eventMiningRanking/fetchEventMiningRanking", fetchEventMiningRanking)
}
