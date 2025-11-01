package user_trade

import (
	"elichika/client"
	"elichika/enum"
	"elichika/generic"
	"elichika/subsystem/user_event/mining"
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

func GetTradeTypeByProductId(session *userdata.Session, productId int32) int32 {
	product := session.Gamedata.TradeProduct[productId]
	trade := session.Gamedata.Trade[product.TradeId]
	if trade != nil {
		return trade.TradeType
	}
	// if trade doesn't exist, then this is event trade, so use normal trade type
	return enum.TradeTypeTrade

}

func GetTrades(session *userdata.Session, tradeType int32) generic.Array[client.Trade] {
	trades := generic.Array[client.Trade]{}
	for _, tradePtr := range session.Gamedata.TradesByType[tradeType] {
		trade := *tradePtr
		for i, product := range trade.Products.Slice {
			product.TradedCount = GetUserTradeProduct(session, product.ProductId)
			trade.Products.Slice[i] = product
		}
		trades.Append(trade)
	}
	event := session.Gamedata.EventActive.GetEventValue()
	if (event != nil) && (event.EventType == enum.EventType1Mining) {
		// add mining event
		tradePtr := session.Gamedata.EventActive.GetEventMining().Trade
		trade := *tradePtr // client.Trade
		for i, product := range trade.Products.Slice {
			product.TradedCount = mining.GetUserEventMiningTradeProduct(session, product.ProductId)
			trade.Products.Slice[i] = product
		}
		trade.StartAt = event.StartAt
		// TODO(extra): in original server, the trade would stay around for 6 days after the event has ended
		// here we couple them and make the trade go away when the event go away:
		// - with a more dynamic trade system, we can decouple the event itself from the event trade, and just spawn in trade whenever necessary
		// - but that require big modification to the trade system, so it's not worth it for now
		// - as a tradeoff, existing currency are not removed whenever an event come back
		trade.EndAt = generic.NewNullable(event.EndAt)
		trades.Append(trade)
	}
	return trades
}
