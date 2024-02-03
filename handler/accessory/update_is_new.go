package accessory

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_accessory"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func updateIsNew(ctx *gin.Context) {
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_accessory.ClearIsNewFlags(session)

	session.Finalize()
	common.JsonResponse(ctx, &response.EmptyResponse{})
}

func init() {
	router.AddHandler("/accessory/updateIsNew", updateIsNew)
}
