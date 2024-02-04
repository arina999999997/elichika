package tower

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func recoveryTowerCardUsedAll(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.RecoveryTowerCardUsedRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.RecoveryTowerCardUsedResponse{
		TowerCardUsedCountRows: session.GetUserTowerCardUsedList(req.TowerId),
		UserModelDiff:          &session.UserModel,
	}
	for i := range resp.TowerCardUsedCountRows.Slice {
		resp.TowerCardUsedCountRows.Slice[i].UsedCount = 0
		resp.TowerCardUsedCountRows.Slice[i].RecoveredCount = 0
		session.UpdateUserTowerCardUsed(req.TowerId, resp.TowerCardUsedCountRows.Slice[i])
	}

	session.Finalize()
	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/tower/recoveryTowerCardUsedAll", recoveryTowerCardUsedAll)
}
