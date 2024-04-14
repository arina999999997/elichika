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

func recoveryTowerCardUsedAll(ctx *gin.Context) {
	req := request.RecoveryTowerCardUsedRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := response.RecoveryTowerCardUsedResponse{
		TowerCardUsedCountRows: user_tower.GetUserTowerCardUsedList(session, req.TowerId),
		UserModelDiff:          &session.UserModel,
	}
	for i := range resp.TowerCardUsedCountRows.Slice {
		resp.TowerCardUsedCountRows.Slice[i].UsedCount = 0
		resp.TowerCardUsedCountRows.Slice[i].RecoveredCount = 0
		user_tower.UpdateUserTowerCardUsed(session, req.TowerId, resp.TowerCardUsedCountRows.Slice[i])
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/tower/recoveryTowerCardUsedAll", recoveryTowerCardUsedAll)
}
