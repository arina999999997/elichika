package client

import (
	"elichika/generic"
)

type EventMarathonUserRanking struct {
	Order           int32                     `json:"order"`
	TotalPoint      generic.Nullable[int32]   `json:"total_point"`
	NextRewardPoint generic.Nullable[int32]   `json:"next_reward_point"`
	RewardContent   generic.Nullable[Content] `json:"reward_content"` // can this be null?
}
