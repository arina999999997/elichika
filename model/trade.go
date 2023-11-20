package model

type TradeProductUser struct {
	UserID      int `xorm:"pk 'user_id'"`
	ProductID   int `xorm:"pk 'product_id'"`
	TradedCount int `xorm:"'traded_count'"`
}

type TradeProduct struct {
	// ProductID is defined to be TradeID * 1000 + some number
	ProductID int `xorm:"pk 'product_id'" json:"product_id"`
	// TradeID inside json is the same as ProductID
	DummyID int `xorm:"-" json:"trade_id"`
	// TradeID inside db is the actual TradeID that this product belong to
	TradeID       int       `xorm:"'trade_id'" json:"-"`
	SourceAmount  int       `json:"source_amount"`         // cost
	StockAmount   *int      `json:"stock_amount"`          // max exchange time, set to null for unlimited
	TradedCount   int       `xorm:"-" json:"traded_count"` // amount traded, store per user
	Contents      []Content `xorm:"-" json:"contents"`     // array but the length is always 1
	ActualContent Content   `xorm:"extends" json:"-" `     // actual content
}

// represent exchange or channel exchange banner
type Trade struct {
	TradeID int `xorm:"pk 'trade_id'" json:"trade_id"`
	// trade type = 1 for normal exchange and 2 for channel exchange
	TradeType       int `xorm:"'trade_type'" json:"-"`
	BannerImagePath struct {
		V string `xorm:"'banner_image_path'" json:"v"`
	} `xorm:"extends" json:"banner_image_path"`
	// items used for exchange
	SourceContentType        int `xorm:"-" json:"source_content_type"`
	SourceContentID          int `xorm:"-" json:"source_content_id"`
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
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_trade_product"] = TradeProductUser{}
}