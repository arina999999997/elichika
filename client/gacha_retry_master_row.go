package client

import (
	"elichika/generic"
)

type GachaRetryMasterRow struct {
	GachaDrawId     int32                   `json:"gacha_draw_id"`
	RetryOrder      int32                   `json:"retry_order"`
	RetryCount      generic.Nullable[int32] `json:"retry_count"`
	PaymentType     generic.Nullable[int32] `json:"payment_type" enum:"GachaPaymentType"`
	PaymentTargetId generic.Nullable[int32] `json:"payment_target_id"`
	PaymentAmount   generic.Nullable[int32] `json:"payment_amount"`
}
