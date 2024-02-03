package info_trigger

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_info_trigger"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func readMemberLoveLevelUp(ctx *gin.Context) {
	// there is no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_info_trigger.ReadAllMemberLoveLevelUpTriggers(session)

	session.Finalize()
	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/infoTrigger/readMemberLoveLevelUp", readMemberLoveLevelUp)
}
