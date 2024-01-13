package client

type Subscription struct {
	SubscriptionMasterId int32 `json:"subscription_master_id"`
	IsTrial              bool  `json:"is_trial"`
}
