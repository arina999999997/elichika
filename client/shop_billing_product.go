package client

import (
	"elichika/generic"
)

type ShopBillingProduct struct {
	ShopProductMasterId int32 `json:"shop_product_master_id"`
	BillingProductType  int32 `json:"billing_product_type" enum:"BillingProductType"` // convert to an enum by the end
	Price               int32 `json:"price"`
	// all of these are not always required, just need to exist for relevant functions
	ShopBillingProductContent       *generic.Array[ShopBillingProductContent] `json:"shop_billing_product_content,omitempty"`
	ShopBillingPlatformProduct      *ShopBillingPlatformProduct               `json:"shop_billing_platform_product,omitempty"` // required for a lot of things, should be the product Id at google's or apple's
	ShopBillingLimitedProductDetail *ShopBillingLimitedProductDetail          `json:"shop_billing_limited_product_detail,omitempty"`
	GiftBox                         *GiftBox                                  `json:"gift_box,omitempty"`
	Subscription                    *Subscription                             `json:"subscription,omitempty"`
}
