package response

import (
	"elichika/client"
)

type FetchSubscriptionPassResponse struct {
	BeforeContinueCount int32 `json:"before_continue_count"`
}

type UpdateSubscriptionResponse struct {
	UserModel        *client.UserModel `json:"user_model"`
	BillingStateInfo BillingStateInfo  `json:"billing_state_info"`
}
