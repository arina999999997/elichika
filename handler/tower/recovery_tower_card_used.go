package tower

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/handler/common"
	"elichika/item"
	"elichika/router"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_tower"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func recoveryTowerCardUsed(ctx *gin.Context) {
	req := request.RecoveryTowerCardUsedRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	for _, cardMasterId := range req.CardMasterIds.Slice {
		cardUsedCount := user_tower.GetUserTowerCardUsed(session, req.TowerId, cardMasterId)
		cardUsedCount.UsedCount--
		cardUsedCount.RecoveredCount++
		user_tower.UpdateUserTowerCardUsed(session, req.TowerId, cardUsedCount)
	}
	// remove the item, this has to be done manually because it involve going back to gems

	has := user_content.GetUserContentByContent(session, item.PerformanceDrink).ContentAmount
	cardCount := int32(req.CardMasterIds.Size())
	if config.Conf.ResourceConfig().ConsumeMiscItems {
		if has >= cardCount {
			user_content.RemoveContent(session, item.PerformanceDrink.Amount(cardCount))
		} else {
			user_content.RemoveContent(session, item.PerformanceDrink.Amount(has))
			user_content.RemoveContent(session, item.StarGem.Amount((cardCount-has)*int32(session.Gamedata.Tower[req.TowerId].RecoverCostBySnsCoin)))
		}
	}

	session.Finalize()
	common.JsonResponse(ctx, &response.RecoveryTowerCardUsedResponse{
		TowerCardUsedCountRows: user_tower.GetUserTowerCardUsedList(session, req.TowerId),
		UserModelDiff:          &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/tower/recoveryTowerCardUsed", recoveryTowerCardUsed)
}
