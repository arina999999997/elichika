package client

import (
	"elichika/generic"
)

type LiveResultActiveEventPointReward struct {
	NextPointReward    generic.Nullable[EventMarathonPointReward] `json:"next_point_reward"`
	GettedPointRewards generic.List[EventMarathonPointReward]     `json:"getted_point_rewards"`
}
