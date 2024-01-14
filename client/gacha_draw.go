package client

import (
	"elichika/generic"
)

type GachaDraw struct {
	GachaDrawMasterId    int32                             `json:"gacha_draw_master_id"`
	RecoverType          int32                             `json:"recover_type" enum:"GachaRecoverType"`
	RecoverAt            generic.Nullable[int64]           `json:"recover_at"`
	DrawCount            int32                             `json:"draw_count"`
	GachaPaymentType     int32                             `json:"gacha_payment_type" enum:"GachaPaymentType"`
	GachaPaymentMasterId int32                             `json:"gacha_payment_master_id"`
	GachaPaymentAmount   int32                             `json:"gacha_payment_amount"`
	GachaPointAmount     generic.Nullable[int32]           `json:"gacha_point_amount"`
	Description          generic.Nullable[LocalizedText]   `json:"description"` // pointer
	IsBonus              bool                              `json:"is_bonus"`
	BonusAppealText      generic.Nullable[LocalizedText]   `json:"bonus_appeal_text"` // pointer
	RetryCount           generic.Nullable[int32]           `json:"retry_count"`
	GachaRetryMasterRows generic.List[GachaRetryMasterRow] `json:"gacha_retry_master_rows"`
	DailyLimit           generic.Nullable[int32]           `json:"daily_limit"`
	DailyInterval        generic.Nullable[int32]           `json:"daily_interval"`
	TermLimit            generic.Nullable[int32]           `json:"term_limit"`
	RemainDayCount       generic.Nullable[int32]           `json:"remain_day_count"`
	RemainTermCount      generic.Nullable[int32]           `json:"remain_term_count"`
	PerformanceId        int32                             `json:"performance_id"`
	IsSubscription       bool                              `json:"is_subscription"`
}
