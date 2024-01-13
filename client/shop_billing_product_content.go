package client

type ShopBillingProductContent struct {
	Amount          int32 `json:"amount"`
	ContentType     int32 `json:"content_type" enum:"ContentType"`
	ContentMasterId int32 `json:"content_master_id"`
	IsPaidContent   bool  `json:"is_paid_content"`
}
