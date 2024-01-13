package client

import (
	"elichika/generic"
)

type BillingStateInfo struct {
	Age                        generic.Nullable[int32] `json:"age"`
	CurrentMonthPurcharsePrice int32                   `json:"current_month_purchase_price"`
}
