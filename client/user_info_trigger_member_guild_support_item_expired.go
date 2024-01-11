package client

type UserInfoTriggerMemberGuildSupportItemExpired struct {
	// always present even if it's not actually expire
	TriggerId int64 `xorm:"pk 'trigger_id'" json:"trigger_id"` // use nano timestamp
	ResetAt   int64 `xorm:"'reset_at'" json:"reset_at"`        // use unit timestamp
}

func (uitmgsie *UserInfoTriggerMemberGuildSupportItemExpired) Id() int64 {
	return uitmgsie.TriggerId
}
