package request

type FetchTradeRequest struct {
	TradeType int32 `json:"trade_type" enum:"TradeType"`
}
