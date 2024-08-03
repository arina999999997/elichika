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

// response: UserModelResponse
func finishEventStory(ctx *gin.Context) {
	req := request.FinishEventStoryRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	common.JsonResponse(ctx, event.FinishEventStory(session, req.StoryEventMasterId, req.IsAutoMode))
}

func init() {
	router.AddHandler("/", "POST", "/eventMarathon/finishEventStory", finishEventStory)
}
