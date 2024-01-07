// response necessary for the subscription feature stuff

package response

type ShopBillingProductContent = int
type ShopBillingLimitedProductDetail = int
type GiftBox = int

type ShopBillingPlatformProduct struct {
	PlatformProductId string `json:"platform_product_id"`
}

type Subscription struct {
	SubscriptionMasterId int  `json:"subscription_master_id"`
	IsTrial              bool `json:"is_trial"`
}

type ShopBillingProduct struct {
	ShopProductMasterId int `json:"shop_product_master_id"`
	BillingProductType  int `json:"billing_product_type"` // convert to an enum by the end
	Price               int `json:"price"`
	// all of these are not always required, just need to exist for relevant functions
	ShopBillingProductContent       []ShopBillingProductContent      `json:"shop_billing_product_content,omitempty"`
	ShopBillingPlatformProduct      *ShopBillingPlatformProduct      `json:"shop_billing_platform_product,omitempty"` // required for a lot of things, should be the product Id at google's or apple's
	ShopBillingLimitedProductDetail *ShopBillingLimitedProductDetail `json:"shop_billing_limited_product_detail,omitempty"`
	GiftBox                         *GiftBox                         `json:"gift_box,omitempty"`
	Subscription                    *Subscription                    `json:"subscription,omitempty"`
}
type BillingStateInfo struct {
	Age                        int `json:"age"`
	CurrentMonthPurcharsePrice int `json:"current_month_purchase_price"`
}
type ShopSubscriptionResponse struct {
	ProductList      []ShopBillingProduct `json:"product_list"`
	BillingStateInfo BillingStateInfo     `json:"billing_state_info"`
}
