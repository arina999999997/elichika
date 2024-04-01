package live

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func surrender(ctx *gin.Context) {
	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	lpDiff := user_live.SurrenderLive(session)

	common.JsonResponse(ctx, &response.SurrenderLiveResponse{
		LpDiff:        lpDiff,
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/live/surrender", surrender)
}
