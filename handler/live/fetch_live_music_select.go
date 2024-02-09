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
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := user_live.FetchLiveMusicSelect(session)

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/live/fetchLiveMusicSelect", fetchLiveMusicSelect)
}
