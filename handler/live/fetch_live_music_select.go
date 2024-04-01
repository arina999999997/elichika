package live

import (
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchLiveMusicSelect(ctx *gin.Context) {
	// ther is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_live.FetchLiveMusicSelect(session)

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/live/fetchLiveMusicSelect", fetchLiveMusicSelect)
}
