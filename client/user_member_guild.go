package client

type UserMemberGuild struct {
	MemberGuildId            int32 `xorm:"pk 'member_guild_id'" json:"member_guild_id"`
	MemberMasterId           int32 `xorm:"'member_master_id'" json:"member_master_id"`
	TotalPoint               int32 `xorm:"'total_point'" json:"total_point"`
	SupportPoint             int32 `xorm:"'support_point'" json:"support_point"`
	LovePoint                int32 `xorm:"'love_point'" json:"love_point"`
	VoltagePoint             int32 `xorm:"'voltage_point'" json:"voltage_point"`
	DailySupportPoint        int32 `xorm:"'daily_support_point'" json:"daily_support_point"`
	DailySupportPointResetAt int32 `xorm:"'daily_support_point_reset_at'" json:"daily_support_point_reset_at"`
	DailyLovePoint           int32 `xorm:"'daily_love_point'" json:"daily_love_point"`
	DailyLovePointResetAt    int32 `xorm:"'daily_love_point_reset_at'" json:"daily_love_point_reset_at"`
	MaxVoltage               int64 `xorm:"'max_voltage'" json:"max_voltage"`
	SupportPointCountResetAt int64 `xorm:"'support_point_count_reset_at'" json:"support_point_count_reset_at"`
	// DailyLovePointResetAt and MaxVoltage 's types are correct (match client), but maybe they meant to hibe the int64 to the reset at
}
