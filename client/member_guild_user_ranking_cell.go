package client

import (
	"elichika/generic"
)

type MemberGuildUserRankingCell struct {
	Order                          generic.Nullable[int32]        `json:"order"`
	TotalPoint                     int32                          `json:"total_point"`
	MemberGuildUserRankingUserData MemberGuildUserRankingUserData `json:"member_guild_user_ranking_user_data"`
}
