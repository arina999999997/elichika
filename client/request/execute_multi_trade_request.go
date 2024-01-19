package request
import (
	"elichika/generic"
	"elichika/client"
)
type ExecuteMultiTradeRequest struct {
	TradeOrders generic.Array[client.TradeOrder] `json:"trade_orders"`
}