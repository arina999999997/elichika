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

func executeMultiTrade(ctx *gin.Context) {
	req := request.ExecuteMultiTradeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	sentToPresentBox := false
	for _, trade := range req.TradeOrders.Slice {
		if user_trade.ExecuteTrade(session, trade.ProductId, trade.TradeCount) {
			sentToPresentBox = true
		}
	}
	sentToPresentBox = sentToPresentBox || (len(session.UnreceivedContent) > 0)

	session.Finalize()
	common.JsonResponse(ctx, response.ExecuteTradeResponse{
		Trades:           user_trade.GetTrades(session, session.Gamedata.Trade[session.Gamedata.TradeProduct[req.TradeOrders.Slice[0].ProductId].TradeId].TradeType),
		IsSendPresentBox: sentToPresentBox,
		UserModelDiff:    &session.UserModel,
	})
}

func init() {
	router.AddHandler("/", "POST", "/trade/executeMultiTrade", executeMultiTrade)
}
