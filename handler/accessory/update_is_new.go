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
	user_accessory.ClearIsNewFlags(ctx.MustGet("session").(*userdata.Session))

	common.JsonResponse(ctx, &response.EmptyResponse{})
}

func init() {
	router.AddHandler("/", "POST", "/accessory/updateIsNew", updateIsNew)
}
