package tower

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_tower"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func clearedTowerFloor(ctx *gin.Context) {
	req := request.ClearedTowerFloorRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	userTower := user_tower.GetUserTower(session, req.TowerId)
	if userTower.ClearedFloor < req.FloorNo {
		userTower.ClearedFloor = req.FloorNo
		user_tower.UpdateUserTower(session, userTower)
	}
	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}

	common.JsonResponse(ctx, &response.ClearedTowerFloorResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/tower/clearedTowerFloor", clearedTowerFloor) // dlp story
}
