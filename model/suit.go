package model

import (
	"elichika/generic"
)

type UserSuit struct {
	SuitMasterId int  `xorm:"pk 'suit_master_id'" json:"suit_master_id"`
	IsNew        bool `xorm:"'is_new'" json:"is_new"`
}

func (us *UserSuit) Id() int64 {
	return int64(us.SuitMasterId)
}
func init() {

	TableNameToInterface["u_suit"] = generic.UserIdWrapper[UserSuit]{}
}
