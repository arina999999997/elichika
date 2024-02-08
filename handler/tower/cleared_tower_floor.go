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
	"github.com/tidwall/gjson"
)

func clearedTowerFloor(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ClearedTowerFloorRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	userTower := user_tower.GetUserTower(session, req.TowerId)
	if userTower.ClearedFloor < req.FloorNo {
		userTower.ClearedFloor = req.FloorNo
		user_tower.UpdateUserTower(session, userTower)
	}
	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}

	session.Finalize()
	common.JsonResponse(ctx, &response.ClearedTowerFloorResponse{
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/tower/clearedTowerFloor", clearedTowerFloor) // dlp story
}
