package client

type TradeOrder struct {
	ProductId int32 `json:"product_id"`
	TradeCount int32 `json:"trade_count"`
}