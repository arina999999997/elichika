package client

type BillingBalanceHistory struct {
	ShopProductId      int32 `json:"shop_product_id"`
	BillingProductType int32 `json:"billing_product_type" enum:"BillingProductType"`
	Amount             int32 `json:"amount"`
}
