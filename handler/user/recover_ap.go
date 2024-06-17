package user

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/handler/common"
	"elichika/item"
	"elichika/router"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_status"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func recoverAp(ctx *gin.Context) {
	req := request.RecoverAPRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	// TODO(hardcoded): Technically items can recover multiple ap at once per item, use use multiple items just to recover 1 AP
	errorResponse := user_status.AddUserAp(session, req.Count.Value)
	if errorResponse != nil {
		common.AlternativeJsonResponse(ctx, errorResponse)
		return
	}

	if config.Conf.ResourceConfig().ConsumeMiscItems {
		switch req.ContentId {
		case item.TrainingTicket.ContentId:
			user_content.RemoveContent(session, item.TrainingTicket.Amount(req.Count.Value))
		case item.StarGem.ContentId:
			if session.UserStatus.ActivityPointPaymentRecoveryDailyResetAt <= session.Time.Unix() {
				session.UserStatus.ActivityPointPaymentRecoveryDailyCount = 0
				session.UserStatus.ActivityPointPaymentRecoveryDailyResetAt = utils.BeginOfNextDay(session.Time).Unix()
			}
			newCount := session.UserStatus.ActivityPointPaymentRecoveryDailyCount + req.Count.Value
			oldCount := session.UserStatus.ActivityPointPaymentRecoveryDailyCount
			session.UserStatus.ActivityPointPaymentRecoveryDailyCount = newCount
			cost := session.Gamedata.ActivityPointRecoveryPrice[newCount].Amount -
				session.Gamedata.ActivityPointRecoveryPrice[oldCount].Amount
			user_content.RemoveContent(session, item.StarGem.Amount(cost))
		}
	}
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/user/recoverAp", recoverAp)
}
