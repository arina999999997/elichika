package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchTradeResponse struct {
	Trades generic.Array[client.Trade] `json:"trades"` // the name is actually _Trades, for some reason
}
