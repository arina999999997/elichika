package client

type UserInfoTriggerMemberGuildRankingShowResultRow struct {
	TriggerId     int64 `json:"trigger_id"`
	MemberGuildId int32 `json:"member_guild_id"`
	ResultAt      int64 `json:"result_at"`
	EndAt         int64 `json:"end_at"`
}
