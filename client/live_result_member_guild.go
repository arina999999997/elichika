package client

type LiveResultMemberGuild struct {
	MemberGuildId       int32 `json:"member_guild_id"`
	ReceiveLovePoint    int32 `json:"receive_love_point"`
	ReceiveVoltagePoint int32 `json:"receive_voltage_point"`
	TotalPoint          int32 `json:"total_point"`
}
