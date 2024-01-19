package client

import (
	"elichika/generic"
)

type ShopTopIsOpen struct {
	ShopEventItemExchangeType generic.Nullable[int32] `json:"shop_event_item_exchange_type" enum:"ShopEventItemExchangeType"`
	IsOpen                    bool                    `json:"is_open"`
}
