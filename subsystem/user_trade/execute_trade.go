package user_trade

import (
	"elichika/client"
	"elichika/enum"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_present"
	"elichika/userdata"
)

// return whether the item is added to present box
func ExecuteTrade(session *userdata.Session, productId, tradeCount int32) bool {
	// update count
	tradedCount := GetUserTradeProduct(session, productId)
	tradedCount += tradeCount
	SetUserTradeProduct(session, productId, tradedCount)

	// award items and take away source item
	product := session.Gamedata.TradeProduct[productId]
	trade := session.Gamedata.Trade[product.TradeId]
	inPresentBox := false
	for _, content := range product.Contents.Slice {
		if content.ContentType == enum.ContentTypeCard {
			tradeCount = 1
			user_present.AddPresent(session, client.PresentItem{
				PresentRouteType: enum.PresentRouteTypeTrade,
				Content:          content.Amount(1),
			})
			inPresentBox = true
		} else {
			user_content.AddContent(session, content.Amount(tradedCount))
		}
	}
	user_content.RemoveContent(session, client.Content{
		ContentType:   trade.SourceContentType,
		ContentId:     trade.SourceContentId,
		ContentAmount: product.SourceAmount * tradeCount,
	})
	return inPresentBox
}
