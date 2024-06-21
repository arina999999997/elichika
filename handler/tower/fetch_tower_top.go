package tower

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_tower"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func fetchTowerTop(ctx *gin.Context) {
	req := request.FetchTowerTopRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	success, failure := user_tower.FetchTowerTop(session, req.TowerId)
	if success != nil {
		common.JsonResponse(ctx, success)
	} else {
		common.AlternativeJsonResponse(ctx, failure)
	}
}

func init() {
	router.AddHandler("/", "POST", "/tower/fetchTowerTop", fetchTowerTop)
}
