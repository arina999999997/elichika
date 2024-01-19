package request

import (
	"elichika/client"
	"elichika/generic"
)

type ExecuteMultiTradeRequest struct {
	TradeOrders generic.Array[client.TradeOrder] `json:"trade_orders"`
}
