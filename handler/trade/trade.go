package trade

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func FetchTrade(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchTradeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	common.JsonResponse(ctx, response.FetchTradeResponse{
		Trades: session.GetTrades(req.TradeType),
	})
}

func ExecuteTrade(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ExecuteTradeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// this only decide whether there's a text saying that things were sent to present box
	sentToPresentBox := session.ExecuteTrade(req.ProductId, req.TradeCount)
	session.Finalize()

	common.JsonResponse(ctx, response.ExecuteTradeResponse{
		Trades:           session.GetTrades(session.Gamedata.Trade[session.Gamedata.TradeProduct[req.ProductId].TradeId].TradeType),
		IsSendPresentBox: sentToPresentBox,
		UserModelDiff:    &session.UserModel,
	})
}

func ExecuteMultiTrade(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ExecuteMultiTradeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	sentToPresentBox := false
	for _, trade := range req.TradeOrders.Slice {
		if session.ExecuteTrade(trade.ProductId, trade.TradeCount) {
			sentToPresentBox = true
		}
	}
	session.Finalize()
	common.JsonResponse(ctx, response.ExecuteTradeResponse{
		Trades:           session.GetTrades(session.Gamedata.Trade[session.Gamedata.TradeProduct[req.TradeOrders.Slice[0].ProductId].TradeId].TradeType),
		IsSendPresentBox: sentToPresentBox,
		UserModelDiff:    &session.UserModel,
	})
}

func init() {
	router.AddHandler("/trade/fetchTrade", FetchTrade)
	router.AddHandler("/trade/executeTrade", ExecuteTrade)
	router.AddHandler("/trade/executeMultiTrade", ExecuteMultiTrade)

}
