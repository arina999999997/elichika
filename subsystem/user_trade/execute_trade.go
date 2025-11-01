package user_trade

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/log"
	"elichika/subsystem/user_content"
	"elichika/subsystem/user_event/mining"
	"elichika/subsystem/user_present"
	"elichika/userdata"
)

// return whether the item is added to present box
func ExecuteTrade(session *userdata.Session, productId, tradeCount int32) bool {
	product := session.Gamedata.TradeProduct[productId]
	trade := session.Gamedata.Trade[product.TradeId]
	// first modify the tracking amount
	var totalTradeCount int32
	var sourceContentType int32
	var sourceContentId int32
	if trade != nil {
		// normal trade
		// update count
		totalTradeCount = GetUserTradeProduct(session, productId)
		totalTradeCount += tradeCount
		SetUserTradeProduct(session, productId, totalTradeCount)
		sourceContentType = trade.SourceContentType
		sourceContentId = trade.SourceContentId
	} else {
		// special trade, this should mean event trade
		event := session.Gamedata.EventActive.GetEventValue()
		if event == nil {
			log.Panic("Invalid trade: trade doesn't exist and there's no event")
		}
		switch event.EventType {
		case enum.EventType1Mining:
			trade := session.Gamedata.EventMining[product.TradeId].Trade
			sourceContentType = trade.SourceContentType
			sourceContentId = trade.SourceContentId
			totalTradeCount = mining.GetUserEventMiningTradeProduct(session, productId)
			totalTradeCount += tradeCount
			mining.SetUserEventMiningTradeProduct(session, productId, totalTradeCount)
		default:
			log.Panic("Event type doesn't support trade")
		}
	}

	// award items and take away source item
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
			user_content.AddContent(session, content.Amount(tradeCount))
		}
	}
	if config.Conf.ResourceConfig().ConsumeExchangeCurrency {
		user_content.RemoveContent(session, client.Content{
			ContentType:   sourceContentType,
			ContentId:     sourceContentId,
			ContentAmount: product.SourceAmount * tradeCount,
		})
	}
	return inPresentBox
}
