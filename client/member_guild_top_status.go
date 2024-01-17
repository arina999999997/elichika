package client

import (
	"elichika/generic"
)

type MemberGuildTopStatus struct {
	MemberGuildRankingAnimationInfo       generic.List[MemberGuildRankingAnimationInfo] `json:"member_guild_ranking_animation_info"`
	MemberGuildRankingResultAnimationInfo generic.List[MemberGuildRankingAnimationInfo] `json:"member_guild_ranking_result_animation_info"`
	IsTopRankingDisplay                   bool                                          `json:"is_top_ranking_display"`
	MemberGuildInfo                       MemberGuildTopInfo                            `json:"member_guild_info"`
}
