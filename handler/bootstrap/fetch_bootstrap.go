package bootstrap

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_bootstrap"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func fetchBootstrap(ctx *gin.Context) {
	req := request.FetchBootstrapRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := user_bootstrap.FetchBootstrap(session, req)

	session.Finalize()
	common.JsonResponse(ctx, resp)
}

func init() {
	router.AddHandler("/bootstrap/fetchBootstrap", fetchBootstrap)
}
