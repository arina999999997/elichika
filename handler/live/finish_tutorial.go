package live

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func finishTutorial(ctx *gin.Context) {
	req := request.FinishLiveRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_live.FinishTutorial(session, req)

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/live/finishTutorial", finishTutorial)
}
