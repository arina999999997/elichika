package mining

import (
	"elichika/userdata"
	"elichika/utils"
)

func GetUserEventMiningTradeProduct(session *userdata.Session, productId int32) int32 {
	result := int32(0)
	exist, err := session.Db.Table("u_event_mining_trade_product").
		Where("user_id = ? AND product_id = ?", session.UserId, productId).
		Cols("traded_count").Get(&result)
	utils.CheckErr(err)
	if !exist {
		result = 0
	}
	return result
}
