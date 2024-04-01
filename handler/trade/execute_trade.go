package trade

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_trade"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func executeTrade(ctx *gin.Context) {
	req := request.ExecuteTradeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// this only decide whether there's a text saying that things were sent to present box
	sentToPresentBox := user_trade.ExecuteTrade(session, req.ProductId, req.TradeCount)
	sentToPresentBox = sentToPresentBox || (len(session.UnreceivedContent) > 0)
	session.Finalize()

	common.JsonResponse(ctx, response.ExecuteTradeResponse{
		Trades:           user_trade.GetTrades(session, session.Gamedata.Trade[session.Gamedata.TradeProduct[req.ProductId].TradeId].TradeType),
		IsSendPresentBox: sentToPresentBox,
		UserModelDiff:    &session.UserModel,
	})
}

func init() {
	router.AddHandler("/trade/executeTrade", executeTrade)
}
