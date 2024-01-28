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
	"github.com/tidwall/gjson"
)

func FetchTrade(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FetchTradeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	common.JsonResponse(ctx, response.FetchTradeResponse{
		Trades: user_trade.GetTrades(session, req.TradeType),
	})
}

func ExecuteTrade(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ExecuteTradeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// this only decide whether there's a text saying that things were sent to present box
	sentToPresentBox := user_trade.ExecuteTrade(session, req.ProductId, req.TradeCount)
	session.Finalize()

	common.JsonResponse(ctx, response.ExecuteTradeResponse{
		Trades:           user_trade.GetTrades(session, session.Gamedata.Trade[session.Gamedata.TradeProduct[req.ProductId].TradeId].TradeType),
		IsSendPresentBox: sentToPresentBox,
		UserModelDiff:    &session.UserModel,
	})
}

func ExecuteMultiTrade(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ExecuteMultiTradeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	sentToPresentBox := false
	for _, trade := range req.TradeOrders.Slice {
		if user_trade.ExecuteTrade(session, trade.ProductId, trade.TradeCount) {
			sentToPresentBox = true
		}
	}
	session.Finalize()
	common.JsonResponse(ctx, response.ExecuteTradeResponse{
		Trades:           user_trade.GetTrades(session, session.Gamedata.Trade[session.Gamedata.TradeProduct[req.TradeOrders.Slice[0].ProductId].TradeId].TradeType),
		IsSendPresentBox: sentToPresentBox,
		UserModelDiff:    &session.UserModel,
	})
}

func init() {
	router.AddHandler("/trade/fetchTrade", FetchTrade)
	router.AddHandler("/trade/executeTrade", ExecuteTrade)
	router.AddHandler("/trade/executeMultiTrade", ExecuteMultiTrade)

}
