package database

import (
	"elichika/generic"
)

type UserLivePartner struct {
	LivePartnerCategoryId int32 `xorm:"pk"`
	CardMasterId          int32
}

func init() {
	AddTable("u_live_partner", generic.UserIdWrapper[UserLivePartner]{})
}
