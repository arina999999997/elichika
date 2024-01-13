package client

type GiftBoxContent struct {
	Amount          int32 `json:"amount"`
	ContentType     int32 `json:"content_type" enum:"ContentType"`
	ContentMasterId int32 `json:"content_master_id"`
	Day             int32 `json:"day"`
}
