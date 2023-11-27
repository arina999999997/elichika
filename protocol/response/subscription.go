package response

type FetchSubscriptionPassResponse struct {
	BeforeContinueCount int `json:"before_continue_count"`
}

type UpdateSubscriptionResponse struct {
	UserModel        *UserModel       `json:"user_model"`
	BillingStateInfo BillingStateInfo `json:"billing_state_info"`
}
