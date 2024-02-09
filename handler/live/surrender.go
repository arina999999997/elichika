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
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	lpDiff := user_live.SurrenderLive(session)

	session.Finalize()
	common.JsonResponse(ctx, &response.SurrenderLiveResponse{
		LpDiff:        lpDiff,
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/live/surrender", surrender)
}
