package client

import (
	"elichika/generic"
)

type BillingConsumeHistory struct {
	Reason               int32                   `json:"reason" enum:"VirtualMoneyConsumeReason"`
	ReasonId             generic.Nullable[int32] `json:"reason_id"`
	VirtualMoneyMasterId int32                   `json:"virtual_money_master_id"`
	PaidMoney            int32                   `json:"paid_money"`
	FreeMoney            int32                   `json:"free_money"`
	ExecutedAt           int64                   `json:"executed_at"`
	NameParam            LocalizedText           `json:"name_param"` // maybe this need to be nullable too, not we're not handling these things just yet
}
