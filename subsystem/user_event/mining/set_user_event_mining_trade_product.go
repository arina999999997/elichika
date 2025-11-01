package mining

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func SetUserEventMiningTradeProduct(session *userdata.Session, productId, newTradedCount int32) {
	record := database.UserTradeProduct{
		ProductId:   productId,
		TradedCount: newTradedCount,
	}
	exist, err := session.Db.Table("u_event_mining_trade_product").
		Where("user_id = ? AND product_id = ?", session.UserId, productId).
		Update(record)
	utils.CheckErr(err)
	if exist == 0 {
		userdata.GenericDatabaseInsert(session, "u_event_mining_trade_product", record)
	}
}
