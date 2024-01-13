package client

import (
	"elichika/generic"
)

type LoginBonusBirthDayMember struct {
	MemberMasterId generic.Nullable[int32] `xorm:"json 'member_master_id'" json:"member_master_id"`
	SuitMasterId   generic.Nullable[int32] `xorm:"json 'suit_master_id'" json:"suit_master_id"`
}
