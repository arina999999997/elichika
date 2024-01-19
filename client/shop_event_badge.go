package client

import (
	"elichika/generic"
)

type ShopEventBadge struct {
	ShopType         int32                   `json:"shop_type" enum:""`
	EventType        generic.Nullable[int32] `json:"event_type" enum:"ShopEventItemExchangeType"`
	ExpiredAt        int64                   `json:"expired_at"`
	Description      LocalizedText           `json:"description"`
	ButtonImageAsset TextureStruktur         `json:"button_image_asset"`
}
