package client

import (
	"elichika/generic"
)

type InPeriodGiftBox struct {
	ShopProductMasterId int32                         `json:"shop_product_master_id"`
	Day                 int32                         `json:"day"`
	PurchasedAt         int64                         `json:"purchased_at"`
	NextAcquireContent  generic.Array[GiftBoxContent] `json:"next_acquire_content"`
	NoticeId            int32                         `json:"notice_id"`
}
