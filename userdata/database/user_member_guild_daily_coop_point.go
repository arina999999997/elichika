package database

type UserMemberGuildDailyCoopPoint struct {
	MemberMasterId int32 `xorm:"pk 'member_master_id'"`
	ResetAt        int64 `xorm:"pk 'reset_at'"`
	TotalPoint     int32 `xorm:"'total_point'"`
}

func init() {
	AddTable("u_member_guild_daily_coop_point", UserMemberGuildDailyCoopPoint{})
}
