package response

import (
	"elichika/generic"
)

type FetchSubscriptionPassResponse struct {
	BeforeContinueCount generic.Nullable[int32] `json:"before_continue_count"`
}
