package tutorial

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_tutorial"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func corePlayableEnd(ctx *gin.Context) {
	// there's no request body
	session := ctx.MustGet("session").(*userdata.Session)

	user_tutorial.CorePlayableEnd(session)

	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/tutorial/corePlayableEnd", corePlayableEnd)
}
