package serverdata

import (
	"elichika/model"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"xorm.io/xorm"
)

func InsertTrade(session *xorm.Session, args []string) {
	// insert trades from json

	file := args[0]
	// tradeType := int(args[1][0]) - '0'

	trades := []model.Trade{}
	err := json.Unmarshal([]byte(utils.ReadAllText(file)), &trades)
	utils.CheckErr(err)

	for i, trade := range trades {
		// trades[i].TradeType = tradeType
		// client use int32 so settle with this until patching client
		trades[i].EndAt = 0x7fffffff
		trades[i].ResetAt = 0x7fffffff
		for j, product := range trade.Products {
			product.TradeId = trade.TradeId
			product.ActualContent = product.Contents[0]
			product.StockAmount = nil // set the stock to inf
			trades[i].Products[j] = product
		}
		_, err = session.Table("s_trade_product").Insert(trades[i].Products)
		utils.CheckErr(err)
	}
	_, err = session.Table("s_trade").Insert(trades)
	utils.CheckErr(err)
}

func TradeCli(session *xorm.Session, args []string) {
	if len(args) == 0 {
		fmt.Println("Invalid params:", args)
		return
	}
	switch args[0] {
	case "insert":
		InsertTrade(session, args[1:])
	default:
		fmt.Println("Invalid params:", args)
		return
	}
}
