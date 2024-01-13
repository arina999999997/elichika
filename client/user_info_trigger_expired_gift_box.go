package client

type UserInfoTriggerExpiredGiftBox struct {
	TotalDays          int32              `json:"total_days"`
	ShopBillingProduct ShopBillingProduct `json:"shop_billing_product"`
	IsRepurchase       bool               `json:"is_repurchase"`
}
