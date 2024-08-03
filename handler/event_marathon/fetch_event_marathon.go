package event_marathon

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/event"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// response: FetchEventMarathonResponse
// alternative response: RecoverableExceptionResponse

func fetchEventMarathon(ctx *gin.Context) {
	req := request.FetchEventMarathonRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	success, failure := event.FetchEventMarathon(session, req.EventId)
	if success != nil {
		common.JsonResponse(ctx, success)
	} else {
		common.AlternativeJsonResponse(ctx, failure)
	}
}

func init() {
	router.AddHandler("/", "POST", "/eventMarathon/fetchEventMarathon", fetchEventMarathon)
}
