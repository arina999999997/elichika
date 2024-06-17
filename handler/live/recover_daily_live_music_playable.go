package live

import (
	"elichika/client/request"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_live"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// response: RecoverDailyLiveMusicPlayableResponse
// alternative respnose: RecoverableExceptionResponse
func recoverDailyLiveMusicPlayable(ctx *gin.Context) {
	req := request.RecoverDailyLiveMusicPlayableRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	fmt.Println(req)

	session := ctx.MustGet("session").(*userdata.Session)

	successResponse, failureResponse := user_live.RecoverDailyLiveMusicPlayable(session, req.LiveId)
	if successResponse != nil {
		common.JsonResponse(ctx, successResponse)
	} else {
		common.AlternativeJsonResponse(ctx, failureResponse)
	}
}

func init() {
	router.AddHandler("/", "POST", "/live/recoverDailyLiveMusicPlayable", recoverDailyLiveMusicPlayable)
}
