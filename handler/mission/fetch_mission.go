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
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := user_mission.FetchMission(session)

	session.Finalize() // fetch mission can trigger daily/weekly mission to reset
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/mission/fetchMission", fetchMission)
}
