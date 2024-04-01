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

func getVoltageRankingDeck(ctx *gin.Context) {
	req := request.GetVoltageRankingDeckRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	common.JsonResponse(ctx, voltage_ranking.GetVoltageRankingDeckResponse(session, req.LiveDifficultyId, req.UserId))
}

func init() {
	router.AddHandler("voltageRanking/getVoltageRankingDeck", getVoltageRankingDeck)
}
