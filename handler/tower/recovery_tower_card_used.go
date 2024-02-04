package tower

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/item"
	"elichika/router"
	"elichika/subsystem/user_content"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func recoveryTowerCardUsed(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.RecoveryTowerCardUsedRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	for _, cardMasterId := range req.CardMasterIds.Slice {
		cardUsedCount := session.GetUserTowerCardUsed(req.TowerId, cardMasterId)
		cardUsedCount.UsedCount--
		cardUsedCount.RecoveredCount++
		session.UpdateUserTowerCardUsed(req.TowerId, cardUsedCount)
	}
	// remove the item, this has to be done manually because it involve going back to gems

	has := user_content.GetUserContentByContent(session, item.PerformanceDrink).ContentAmount
	cardCount := int32(req.CardMasterIds.Size())
	if has >= cardCount {
		user_content.RemoveContent(session, item.PerformanceDrink.Amount(cardCount))
	} else {
		user_content.RemoveContent(session, item.PerformanceDrink.Amount(has))
		user_content.RemoveContent(session, item.StarGem.Amount((cardCount-has)*int32(session.Gamedata.Tower[req.TowerId].RecoverCostBySnsCoin)))
	}

	session.Finalize()
	common.JsonResponse(ctx, &response.RecoveryTowerCardUsedResponse{
		TowerCardUsedCountRows: session.GetUserTowerCardUsedList(req.TowerId),
		UserModelDiff:          &session.UserModel,
	})
}

func init() {
	router.AddHandler("/tower/recoveryTowerCardUsed", recoveryTowerCardUsed)
}
