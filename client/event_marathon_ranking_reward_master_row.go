package client

import (
	"elichika/generic"
)

type EventMarathonRankingRewardMasterRow struct {
	RankingRewardMasterId  int32                   `xorm:"'ranking_reward_master_id'" json:"ranking_reward_master_id"`
	UpperRank              int32                   `xorm:"'upper_rank'" json:"upper_rank"`
	LowerRank              generic.Nullable[int32] `xorm:"json 'lower_rank'" json:"lower_rank"`
	RewardGroupId          int32                   `xorm:"'reward_group_id'" json:"reward_group_id"`
	RankingResultPrizeType int32                   `xorm:"'ranking_result_prize_type'" json:"ranking_result_prize_type"`
}
