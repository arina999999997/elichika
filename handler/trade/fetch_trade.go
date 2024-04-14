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

func fetchTrade(ctx *gin.Context) {
	req := request.FetchTradeRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	common.JsonResponse(ctx, response.FetchTradeResponse{
		Trades: user_trade.GetTrades(session, req.TradeType),
	})
}

func init() {
	router.AddHandler("/", "POST", "/trade/fetchTrade", fetchTrade)
}
