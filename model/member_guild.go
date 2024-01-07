package model

import (
	"elichika/generic"
)

type UserMemberGuild struct {
	MemberGuildId            int `xorm:"pk 'member_guild_id'" json:"member_guild_id"`
	MemberMasterId           int `xorm:"'member_master_id'" json:"member_master_id"`
	TotalPoint               int `xorm:"'total_point'" json:"total_point"`
	SupportPoint             int `xorm:"'support_point'" json:"support_point"`
	LovePoint                int `xorm:"'love_point'" json:"love_point"`
	VoltagePoint             int `xorm:"'voltage_point'" json:"voltage_point"`
	DailySupportPoint        int `xorm:"'daily_support_point'" json:"daily_support_point"`
	DailySupportPointResetAt int `xorm:"'daily_support_point_reset_at'" json:"daily_support_point_reset_at"`
	DailyLovePoint           int `xorm:"'daily_love_point'" json:"daily_love_point"`
	DailyLovePointResetAt    int `xorm:"'daily_love_point_reset_at'" json:"daily_love_point_reset_at"`
	MaxVoltage               int `xorm:"'max_voltage'" json:"max_voltage"`
	SupportPointCountResetAt int `xorm:"'support_point_count_reset_at'" json:"support_point_count_reset_at"`
}

func (umg *UserMemberGuild) Id() int64 {
	return int64(umg.MemberGuildId)
}

type UserMemberGuildSupportItem struct {
	SupportItemId      int   `xorm:"'support_item_id'" json:"support_item_id"`
	Amount             int64 `xorm:"'amount'" json:"amount"`
	SupportItemResetAt int   `xorm:"'support_item_reset_at'" json:"support_item_reset_at"`
}

func (umgsi *UserMemberGuildSupportItem) Id() int64 {
	return int64(umgsi.SupportItemId)
}

func init() {

	TableNameToInterface["u_member_guild"] = generic.UserIdWrapper[UserMemberGuild]{}
	TableNameToInterface["u_member_guild_support_item"] = generic.UserIdWrapper[UserMemberGuildSupportItem]{}
}
