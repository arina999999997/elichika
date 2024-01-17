package client

import (
	"elichika/generic"
)

type MemberGuildRankingOneTerm struct {
	MemberGuildId int32                                       `json:"member_guild_id"`
	StartAt       int64                                       `json:"start_at"`
	EndAt         int64                                       `json:"end_at"`
	Channels      generic.List[MemberGuildRankingOneTermCell] `json:"channels"`
}
