package handler

import (
	"elichika/config"
	"elichika/gamedata"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// TODO(refactor): Change to use request and response types
func FetchTrade(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type FetchTradeReq struct {
		TradeType int32 `json:"trade_type"`
	}
	req := FetchTradeReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	trades := session.GetTrades(req.TradeType)

	signBody, _ := sjson.Set("{}", "trades", trades)
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func ExecuteTrade(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type ExecuteTradeReq struct {
		ProductId  int `json:"product_id"`
		TradeCount int `json:"trade_count"`
	}
	req := ExecuteTradeReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	sentToPresentBox := session.ExecuteTrade(req.ProductId, req.TradeCount)

	signBody := session.Finalize("{}", "user_model_diff")
	// this only decide whether there's a text saying that things were sent to present box
	signBody, _ = sjson.Set(signBody, "is_send_present_box", sentToPresentBox)

	tradeType := gamedata.Trade[gamedata.TradeProduct[req.ProductId].TradeId].TradeType
	signBody, _ = sjson.Set(signBody, "trades", session.GetTrades(tradeType))

	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func ExecuteMultiTrade(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type ExecuteMultiTradeReq struct {
		TradeOrders []struct {
			ProductId  int `json:"product_id"`
			TradeCount int `json:"trade_count"`
		} `json:"trade_orders"`
	}
	req := ExecuteMultiTradeReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	gamedata := ctx.MustGet("gamedata").(*gamedata.Gamedata)

	sentToPresentBox := false
	for _, trade := range req.TradeOrders {
		if session.ExecuteTrade(trade.ProductId, trade.TradeCount) {
			sentToPresentBox = true
		}
	}

	signBody := session.Finalize("{}", "user_model_diff")
	// this only decide whether there's a text saying that things were sent to present box
	signBody, _ = sjson.Set(signBody, "is_send_present_box", sentToPresentBox)

	tradeType := gamedata.Trade[gamedata.TradeProduct[req.TradeOrders[0].ProductId].TradeId].TradeType
	signBody, _ = sjson.Set(signBody, "trades", session.GetTrades(tradeType))

	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)

}
