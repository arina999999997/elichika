package database

import (
	"elichika/generic"
)

type UserTradeProduct struct {
	ProductId   int32 `xorm:"pk 'product_id'"`
	TradedCount int32 `xorm:"'traded_count'"`
}

func init() {
	AddTable("u_trade_product", generic.UserIdWrapper[UserTradeProduct]{})
}
