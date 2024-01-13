package client

import (
	"elichika/generic"
)

type BootstrapSubscription struct {
	ContinueRewards generic.List[SubscriptionContinueReward] `json:"continue_rewards"`
}
