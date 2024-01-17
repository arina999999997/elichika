package client

import (
	"elichika/generic"
)

type MemberGuildTopInfo struct {
	MemberGuildUserRankingOrder int32                   `json:"member_guild_user_ranking_order"`
	MemberGuildUserRankingPoint int32                   `json:"member_guild_user_ranking_point"`
	DailyCoopPoint              int32                   `json:"daily_coop_point"`
	CoopRewardPeriodAt          generic.Nullable[int64] `json:"coop_reward_period_at"`
}
