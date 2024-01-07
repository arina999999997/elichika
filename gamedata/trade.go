// trade or exchange
// trades are stored inside both serverdata.db and masterdata.db
// for the clients to read them, mismatch between serverdata.db and masterdata.db will be resolved toward masterdata.db side
// i.e. the masterdata.db Id will be used and such
// Requirements:
// - TradeId in serverdata.db/s_trade must be the same TradeId in masterdata.db/m_trade
// - TradeId reference in serverdata.db/s_trade_product must refer to the correct TradeId above
// - The number of products per TradeId in serverdata.db/s_trade_product should be the same as the
// amount of products per TradeId in m_trade_product. If they are not the same, some items will be ignored
// or filled with undetermined items

package gamedata

import (
	"elichika/client"
	"elichika/dictionary"
	"elichika/model"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type Trade struct {
	TradesByType [3][]model.Trade    // map from trade type to array of Trade
	Trades       map[int]model.Trade // map from TradeId to Trade
	Products     map[int]model.TradeProduct
}

func loadTrade(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading Trade")
	gamedata.Trade = make(map[int]*model.Trade)
	gamedata.TradeProduct = make(map[int]*model.TradeProduct)
	err := serverdata_db.Table("s_trade").Find(&gamedata.Trade)
	utils.CheckErr(err)

	for id, trade := range gamedata.Trade {
		exist, err := masterdata_db.Table("m_trade").Where("id = ?", id).
			Cols("trade_type", "source_content_type", "source_content_id").Get(
			&trade.TradeType, &trade.SourceContentType, &trade.SourceContentId)
		utils.CheckErr(err)
		if !exist {
			fmt.Println("Warning: Skipped trade ", id, " (did not exist in masterdata.db)")
			delete(gamedata.Trade, id)
			continue
		}
		// server and client product_id might not be the same, we need to sync it here
		serverProducts := []model.TradeProduct{}
		err = serverdata_db.Table("s_trade_product").Where("trade_id = ?", id).
			OrderBy("product_id").Find(&serverProducts)
		utils.CheckErr(err)
		clientProductIds := []int{}
		err = masterdata_db.Table("m_trade_product").Where("trade_master_id = ?", id).
			OrderBy("id").Cols("id").Find(&clientProductIds)
		utils.CheckErr(err)

		n := len(serverProducts)
		m := len(clientProductIds)
		for ; n < m; n++ { // if server have less than necessary append random product
			serverProducts = append(
				serverProducts,
				model.TradeProduct{
					TradeId:      id,
					SourceAmount: 1,
					StockAmount:  nil,
					ActualContent: client.Content{
						ContentType:   10,
						ContentId:     0,
						ContentAmount: 1,
					},
				})
		}
		serverProducts = serverProducts[0:m] // if server have more then reduce to what client have

		for i := 0; i < m; i++ { // need to use client's id
			serverProducts[i].ProductId = clientProductIds[i]
			serverProducts[i].DummyId = clientProductIds[i]
			serverProducts[i].Contents = append(serverProducts[i].Contents, serverProducts[i].ActualContent)
			gamedata.TradeProduct[clientProductIds[i]] = &serverProducts[i]
		}
		trade.Products = serverProducts
		gamedata.TradesByType[trade.TradeType] = append(gamedata.TradesByType[trade.TradeType], trade)
	}
}

func init() {
	addLoadFunc(loadTrade)
}
