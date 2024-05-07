package database

import (
	"elichika/generic"
)

type UserMemberGuildRankingRewardReceived struct {
	MemberGuildId int32 `xorm:"member_guild_id"`
}

func init() {
	AddTable("u_member_guild_ranking_reward_received", generic.UserIdWrapper[UserMemberGuildRankingRewardReceived]{})
}
