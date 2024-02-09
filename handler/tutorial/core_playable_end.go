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
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_tutorial.CorePlayableEnd(session)

	session.Finalize()
	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/tutorial/corePlayableEnd", corePlayableEnd)
}
