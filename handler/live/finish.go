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

func finish(ctx *gin.Context) {
	req := request.FinishLiveRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := user_live.FinishLive(session, req)

	session.Finalize()
	common.JsonResponse(ctx, resp)
}

func init() {
	router.AddHandler("/live/finish", finish)
}
