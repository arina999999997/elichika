package client

type UserMemberGuildSupportItem struct {
	SupportItemId      int32 `xorm:"'support_item_id'" json:"support_item_id"`
	Amount             int32 `xorm:"'amount'" json:"amount"`
	SupportItemResetAt int32 `xorm:"'support_item_reset_at'" json:"support_item_reset_at"`
}

func (umgsi *UserMemberGuildSupportItem) Id() int64 {
	return int64(umgsi.SupportItemId)
}
