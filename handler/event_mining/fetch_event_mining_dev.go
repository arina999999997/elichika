//go:build dev

package event_mining

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/event"
	"elichika/userdata"
	"elichika/utils"
	"elichika/webui/event_mining_dev"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

// response: FetchEventMiningResponse
// alternative response: RecoverableExceptionResponse

func fetchEventMining(ctx *gin.Context) {
	req := request.FetchEventMiningRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		// special case for developer mode
		response := &response.FetchEventMiningResponse{
			EventMiningTopStatus: event_mining_dev.TopStatus,
			UserModelDiff:        &session.UserModel,
		}
		common.JsonResponse(ctx, response)
		return
	}
	success, failure := event.FetchEventMining(session, req.EventId)
	if success != nil {
		common.JsonResponse(ctx, success)
	} else {
		common.AlternativeJsonResponse(ctx, failure)
	}
}

func init() {
	router.AddHandler("/", "POST", "/eventMining/fetchEventMining", fetchEventMining)
}
