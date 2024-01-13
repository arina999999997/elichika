package client

import (
	"elichika/generic"
)

type ShopBillingLimitedProductDetail struct {
	BeginnerTerm                     generic.Nullable[int64]                  `json:"beginner_term"`
	SaleEndAt                        generic.Nullable[int64]                  `json:"sale_end_at"`
	LimitAmount                      generic.Nullable[int32]                  `json:"limit_amount"`
	LimitedRemainingAmount           generic.Nullable[int32]                  `json:"limited_remaining_amount"`
	ParentShopBillingPlatformProduct *ShopBillingPlatformProduct              `json:"parent_shop_billing_platform_product"`
	ParentShopBillingProductContent  generic.Array[ShopBillingProductContent] `json:"parent_shop_billing_product_content"`
}
