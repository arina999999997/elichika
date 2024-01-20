package client

import (
	"elichika/generic"
)

type LiveEventCommonInfo struct {
	EventId             int32                                           `json:"event_id"`
	EventType           int32                                           `json:"event_type"`
	ClosedAt            int64                                           `json:"closed_at"`
	BannerStruktur      TextureStruktur                                 `json:"banner_struktur"`
	PointBoostContentId generic.Nullable[int32]                         `json:"point_boost_content_id"`
	EventMusics         generic.Array[LiveEventMiningVoltageMusicBadge] `json:"event_musics"`
}
