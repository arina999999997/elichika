package user_trade

import (
	"elichika/client"
	"elichika/generic"
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func GetUserTradeProduct(session *userdata.Session, productId int32) int32 {
	result := int32(0)
	exist, err := session.Db.Table("u_trade_product").
		Where("user_id = ? AND product_id = ?", session.UserId, productId).
		Cols("traded_count").Get(&result)
	utils.CheckErr(err)
	if !exist {
		result = 0
	}
	return result
}

func SetUserTradeProduct(session *userdata.Session, productId, newTradedCount int32) {
	record := database.UserTradeProduct{
		ProductId:   productId,
		TradedCount: newTradedCount,
	}
	exist, err := session.Db.Table("u_trade_product").
		Where("user_id = ? AND product_id = ?", session.UserId, productId).
		Update(record)
	utils.CheckErr(err)
	if exist == 0 {
		userdata.GenericDatabaseInsert(session, "u_trade_product", record)
	}
}

func GetTrades(session *userdata.Session, tradeType int32) generic.Array[client.Trade] {
	trades := generic.Array[client.Trade]{}
	for _, trade_ptr := range session.Gamedata.TradesByType[tradeType] {
		trade := *trade_ptr
		for j, product := range trade.Products.Slice {
			product.TradedCount = GetUserTradeProduct(session, product.ProductId)
			trade.Products.Slice[j] = product
		}
		trades.Append(trade)
	}
	return trades
}
