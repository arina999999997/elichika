package client

type BillingDepositHistory struct {
	ShopProductId      int32 `json:"shop_product_id"`
	BillingProductType int32 `json:"billing_product_type" enum:"BillingProductType"`
	Price              int32 `json:"price"`
	ExecutedAt         int64 `json:"executed_at"`
}
