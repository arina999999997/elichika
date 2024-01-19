package model

import (
	"elichika/generic"
)

type TradeProductUser struct {
	ProductId   int `xorm:"pk 'product_id'"`
	TradedCount int `xorm:"'traded_count'"`
}

func init() {

	TableNameToInterface["u_trade_product"] = generic.UserIdWrapper[TradeProductUser]{}
}
