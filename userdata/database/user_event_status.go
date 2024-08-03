package database

import (
	"elichika/generic"
)

// These are some info not stored by user model
type UserEventStatus struct {
	EventId int32 `xorm:"pk 'event_id'"`

	IsFirstAccess bool `xorm:"'is_first_access'"`
	// whether the event top state need to be shown as new.
	IsNew bool `xorm:"'is_new'"`

	IsRewardReceived bool `xorm:"'is_reward_received'"`
}

func init() {
	AddTable("u_event_status", generic.UserIdWrapper[UserEventStatus]{})
}
