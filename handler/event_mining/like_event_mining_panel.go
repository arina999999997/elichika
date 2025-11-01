package event_mining

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

// response: LikeEventMiningPanelResponse
func likeEventMiningPanel(ctx *gin.Context) {
	req := request.LikeEventMiningPanelRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)
	common.JsonResponse(ctx, event.LikeEventMiningPanel(session, req.EventId, req.ThumbnailCellId))
}

func init() {
	router.AddHandler("/", "POST", "/eventMining/likeEventMiningPanel", likeEventMiningPanel)
}
