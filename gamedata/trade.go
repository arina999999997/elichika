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
	"elichika/generic"
	"elichika/item"
	"elichika/utils"

	"fmt"

	"xorm.io/xorm"
)

type Trade struct {
	TradeId                  int32                              `xorm:"pk" json:"trade_id"`
	TradeType                int32                              `xorm:"-" enum:"TradeType"`
	BannerImagePath          client.TextureStruktur             `xorm:"json" json:"banner_image_path"`
	SourceContentType        int32                              `json:"source_content_type" enum:"ContentType"`
	SourceContentId          int32                              `json:"source_content_id"`
	SourceThumbnailAssetPath client.TextureStruktur             `xorm:"json" json:"source_thumbnail_asset_path"`
	StartAt                  int64                              `json:"start_at"`
	EndAt                    generic.Nullable[int64]            `xorm:"json" json:"end_at"`
	ResetAt                  generic.Nullable[int64]            `xorm:"json" json:"reset_at"`
	MonthlyReset             bool                               `json:"monthly_reset"`
	Products                 generic.Array[client.TradeProduct] `xorm:"-" json:"products"`
}

func (t Trade) ToClientTrade() *client.Trade {
	return &client.Trade{
		TradeId:                  t.TradeId,
		BannerImagePath:          t.BannerImagePath,
		SourceContentType:        t.SourceContentType,
		SourceContentId:          t.SourceContentId,
		SourceThumbnailAssetPath: t.SourceThumbnailAssetPath,
		StartAt:                  t.StartAt,
		EndAt:                    t.EndAt,
		ResetAt:                  t.ResetAt,
		MonthlyReset:             t.MonthlyReset,
		Products:                 t.Products,
	}
}

// TODO(trade): Have proper gamedata types
func loadTrade(gamedata *Gamedata, masterdata_db, serverdata_db *xorm.Session, dictionary *dictionary.Dictionary) {
	fmt.Println("Loading Trade")
	gamedata.Trade = make(map[int32]*Trade)
	gamedata.TradeProduct = make(map[int32]*client.TradeProduct)
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
		serverProducts := []client.TradeProduct{}
		err = serverdata_db.Table("s_trade_product").Where("trade_id = ?", id).
			OrderBy("product_id").Find(&serverProducts)
		utils.CheckErr(err)
		clientProductIds := []int32{}
		err = masterdata_db.Table("m_trade_product").Where("trade_master_id = ?", id).
			OrderBy("id").Cols("id").Find(&clientProductIds)
		utils.CheckErr(err)

		n := len(serverProducts)
		m := len(clientProductIds)
		for ; n < m; n++ { // if server have less than necessary append random product
			defaultContents := generic.Array[client.Content]{}
			defaultContents.Append(item.StarGem)
			serverProducts = append(serverProducts, client.TradeProduct{
				TradeId:      id,
				SourceAmount: 1,
				Contents:     defaultContents,
			})
		}
		serverProducts = serverProducts[0:m] // if server have more then reduce to what client have

		for i := 0; i < m; i++ { // need to use client's id
			serverProducts[i].ProductId = clientProductIds[i]
			gamedata.TradeProduct[clientProductIds[i]] = &serverProducts[i]
		}
		trade.Products.Slice = serverProducts
		gamedata.TradesByType[trade.TradeType] = append(gamedata.TradesByType[trade.TradeType], trade.ToClientTrade())
	}
}

func init() {
	addLoadFunc(loadTrade)
}
