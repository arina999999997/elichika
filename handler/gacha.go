package handler

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/gacha"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchGachaMenu(ctx *gin.Context) {
	// there is no request body
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	JsonResponse(ctx, &response.FetchGachaMenuResponse{
		GachaList:     session.GetGachaList(),
		UserModelDiff: &session.UserModel,
	})
}

func GachaDraw(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.DrawGachaRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	if session.UserStatus.TutorialPhase == enum.TutorialPhaseGacha {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseFinal
	}

	ctx.Set("session", session)
	gacha, resultCards := gacha.HandleGacha(ctx, req)

	session.Finalize("{}", "dummy")
	JsonResponse(ctx, response.DrawGachaResponse{
		Gacha:         gacha,
		ResultCards:   resultCards,
		UserModelDiff: &session.UserModel,
	})
}
