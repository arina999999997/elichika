package gacha

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_gacha"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func draw(ctx *gin.Context) {
	req := request.DrawGachaRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseGacha {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseFinal
	}

	ctx.Set("session", session)
	gacha, resultCards := user_gacha.HandleGacha(ctx, req)

	session.Finalize()
	common.JsonResponse(ctx, response.DrawGachaResponse{
		Gacha:         gacha,
		ResultCards:   resultCards,
		UserModelDiff: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/gacha/draw", draw)
}
