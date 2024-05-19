package live_partners

import (
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_social"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetch(ctx *gin.Context) {
	// there's no request body
	session := ctx.MustGet("session").(*userdata.Session)
	common.JsonResponse(ctx, user_social.GetLivePartners(session))
}

func init() {
	router.AddHandler("/", "POST", "/livePartners/fetch", fetch)
}
