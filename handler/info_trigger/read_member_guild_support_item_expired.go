package info_trigger

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_info_trigger"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func readMemberGuildSupportItemExpired(ctx *gin.Context) {
	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	user_info_trigger.ReadMemberGuildSupportItemExpired(session)

	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/infoTrigger/readMemberGuildSupportItemExpired", readMemberGuildSupportItemExpired)
}
