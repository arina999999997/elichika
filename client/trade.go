package client

import (
	"elichika/generic"
)

type Trade struct {
	TradeId                  int32                       `xorm:"pk" json:"trade_id"`
	BannerImagePath          TextureStruktur             `xorm:"json" json:"banner_image_path"`
	SourceContentType        int32                       `json:"source_content_type" enum:"ContentType"`
	SourceContentId          int32                       `json:"source_content_id"`
	SourceThumbnailAssetPath TextureStruktur             `xorm:"json" json:"source_thumbnail_asset_path"`
	StartAt                  int64                       `json:"start_at"`
	EndAt                    generic.Nullable[int64]     `xorm:"json" json:"end_at"`
	ResetAt                  generic.Nullable[int64]     `xorm:"json" json:"reset_at"`
	MonthlyReset             bool                        `json:"monthly_reset"`
	Products                 generic.Array[TradeProduct] `xorm:"-" json:"products"`
}
