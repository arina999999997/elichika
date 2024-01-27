package live

import (
	"elichika/client/response"
	"elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_status"
	"elichika/userdata"
	"elichika/utils"

	"github.com/gin-gonic/gin"
)

func surrender(ctx *gin.Context) {
	// there is no request body

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	exist, _, startReq := session.LoadUserLive()
	utils.MustExist(exist)
	session.ClearUserLive()
	// remove only half the LP
	lpCost := session.Gamedata.LiveDifficulty[startReq.LiveDifficultyId].ConsumedLP / 2
	user_status.AddUserLp(session, -lpCost)

	session.Finalize()
	common.JsonResponse(ctx, &response.SurrenderLiveResponse{
		LpDiff:        generic.NewNullable(lpCost),
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/live/surrender", surrender)
}
