package model

import (
	"elichika/client"
	"elichika/generic"
)

type TradeProductUser struct {
	ProductId   int `xorm:"pk 'product_id'"`
	TradedCount int `xorm:"'traded_count'"`
}

type TradeProduct struct {
	// ProductId is defined to be TradeId * 1000 + some number
	ProductId int `xorm:"pk 'product_id'" json:"product_id"`
	// TradeId inside json is the same as ProductId
	DummyId int `xorm:"-" json:"trade_id"`
	// TradeId inside db is the actual TradeId that this product belong to
	TradeId       int              `xorm:"'trade_id'" json:"-"`
	SourceAmount  int              `json:"source_amount"`         // cost
	StockAmount   *int             `json:"stock_amount"`          // max exchange time, set to null for unlimited
	TradedCount   int              `xorm:"-" json:"traded_count"` // amount traded, store per user
	Contents      []client.Content `xorm:"-" json:"contents"`     // array but the length is always 1
	ActualContent client.Content   `xorm:"extends" json:"-" `     // actual content
}

// represent exchange or channel exchange banner
type Trade struct {
	TradeId int `xorm:"pk 'trade_id'" json:"trade_id"`
	// trade type = 1 for normal exchange and 2 for channel exchange
	TradeType       int32 `xorm:"'trade_type'" json:"-" enum:"TradeType"`
	BannerImagePath struct {
		V string `xorm:"'banner_image_path'" json:"v"`
	} `xorm:"extends" json:"banner_image_path"`
	// items used for exchange
	SourceContentType        int32 `xorm:"-" json:"source_content_type"`
	SourceContentId          int32 `xorm:"-" json:"source_content_id"`
	SourceThumbnailAssetPath struct {
		V string `xorm:"'source_thumbnail_asset_path'" json:"v"`
	} `xorm:"extends" json:"source_thumbnail_asset_path"`
	// unix seconds
	StartAt      int64 `json:"start_at"`
	EndAt        int64 `json:"end_at"`
	ResetAt      int64 `json:"reset_at"`
	MonthlyReset bool  `json:"monthly_reset"`

	Products []TradeProduct `xorm:"-" json:"products"`
}

func init() {

	TableNameToInterface["u_trade_product"] = generic.UserIdWrapper[TradeProductUser]{}
}
