package user

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_status"
	"elichika/userdata"
	"elichika/utils"

	"github.com/gin-gonic/gin"
)

func recoverLpSubscription(ctx *gin.Context) {
	// there is no request body
	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	user_status.AddUserLp(session, session.Gamedata.UserRank[session.UserStatus.Rank].MaxLp)
	session.UserStatus.LivePointSubscriptionRecoveryDailyCount = 1 // 1 mean used
	session.UserStatus.LivePointSubscriptionRecoveryDailyResetAt = utils.StartOfNextDay(session.Time).Unix()

	session.Finalize()
	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/user/recoverLpSubscription", recoverLpSubscription)
}
