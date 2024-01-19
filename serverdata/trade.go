package serverdata

import (
	"elichika/client"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"xorm.io/xorm"
)

func InsertTrade(session *xorm.Session, args []string) {
	// insert trades from json

	file := args[0]

	trades := []client.Trade{}
	err := json.Unmarshal([]byte(utils.ReadAllText(file)), &trades)
	utils.CheckErr(err)

	for i, trade := range trades {
		trades[i].EndAt.HasValue = false
		trades[i].ResetAt.HasValue = false
		for j, product := range trade.Products.Slice {
			product.TradeId = trade.TradeId
			product.StockAmount.HasValue = false // set the stock to inf
			trades[i].Products.Slice[j] = product
		}
		_, err = session.Table("s_trade_product").Insert(trades[i].Products.Slice)
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
