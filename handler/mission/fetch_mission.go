package mission

import (
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_mission"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

func fetchMission(ctx *gin.Context) {
	// there is no request body
	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_mission.FetchMission(session)

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/mission/fetchMission", fetchMission)
}
