package client

import (
	"elichika/generic"
)

type Gacha struct {
	GachaMasterId         int32                             `json:"gacha_master_id"`
	GachaType             int32                             `json:"gacha_type" enum:"GachaType"`
	GachaDrawType         int32                             `json:"gacha_draw_type"`
	Title                 LocalizedText                     `json:"title"`
	BannerImageAsset      TextureStruktur                   `json:"banner_image_asset"`
	IsTimeLimited         bool                              `json:"is_time_limited"`
	EndAt                 generic.Nullable[int64]           `json:"end_at"`
	PointMasterId         generic.Nullable[int32]           `json:"point_master_id"`
	PointExchangeExpireAt generic.Nullable[int64]           `json:"point_exchange_expire_at"`
	AppealAt              int64                             `json:"appeal_at"`
	NoticeId              int32                             `json:"notice_id"`
	AppealView            int32                             `json:"appeal_view"`
	GachaAppeals          generic.List[GachaAppeal]         `json:"gacha_appeals"`
	GachaDraws            generic.List[GachaDraw]           `json:"gacha_draws"`
	GachaDrawStepup       generic.Nullable[GachaDrawStepup] `json:"gacha_draw_stepup"` // this is a pointer but can be nulls
}
