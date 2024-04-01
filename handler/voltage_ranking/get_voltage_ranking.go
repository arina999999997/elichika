package voltage_ranking

import (
	"elichika/client/request"
	// "elichika/client/response"
	// "elichika/enum"
	// "elichika/generic"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/voltage_ranking"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func getVoltageRanking(ctx *gin.Context) {
	req := request.GetVoltageRankingRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	common.JsonResponse(ctx, voltage_ranking.GetVoltageRankingResponse(session, req.LiveDifficultyId))
}

func init() {
	router.AddHandler("voltageRanking/getVoltageRanking", getVoltageRanking)
}
