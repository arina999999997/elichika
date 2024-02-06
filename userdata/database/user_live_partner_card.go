package database

import (
	"elichika/generic"
)

type UserLivePartnerCard struct {
	LivePartnerCategoryId int32 `xorm:"pk"`
	CardMasterId          int32
}

func init() {
	AddTable("u_live_partner_card", generic.UserIdWrapper[UserLivePartnerCard]{})
}
