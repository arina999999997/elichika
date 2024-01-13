package response

import (
	"elichika/client"
)

type UpdateSubscriptionResponse struct {
	UserModel        *client.UserModel       `json:"user_model"`
	BillingStateInfo client.BillingStateInfo `json:"billing_state_info"`
}
