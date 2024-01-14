package client

type RetryGacha struct {
	GachaDrawMasterId int32 `json:"gacha_draw_master_id"`
	RemainRetryCount  int32 `json:"remain_retry_count"`
	ExpireAt          int64 `json:"expire_at"`
}
