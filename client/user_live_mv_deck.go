package client

import (
	"elichika/generic"
)

type UserLiveMvDeck struct {
	LiveMasterId     int32                   `xorm:"pk 'live_master_id'" json:"live_master_id"`
	MemberMasterId1  generic.Nullable[int32] `xorm:"json 'member_master_id_1'" json:"member_master_id_1"`
	MemberMasterId2  generic.Nullable[int32] `xorm:"json 'member_master_id_2'" json:"member_master_id_2"`
	MemberMasterId3  generic.Nullable[int32] `xorm:"json 'member_master_id_3'" json:"member_master_id_3"`
	MemberMasterId4  generic.Nullable[int32] `xorm:"json 'member_master_id_4'" json:"member_master_id_4"`
	MemberMasterId5  generic.Nullable[int32] `xorm:"json 'member_master_id_5'" json:"member_master_id_5"`
	MemberMasterId6  generic.Nullable[int32] `xorm:"json 'member_master_id_6'" json:"member_master_id_6"`
	MemberMasterId7  generic.Nullable[int32] `xorm:"json 'member_master_id_7'" json:"member_master_id_7"`
	MemberMasterId8  generic.Nullable[int32] `xorm:"json 'member_master_id_8'" json:"member_master_id_8"`
	MemberMasterId9  generic.Nullable[int32] `xorm:"json 'member_master_id_9'" json:"member_master_id_9"`
	MemberMasterId10 generic.Nullable[int32] `xorm:"json 'member_master_id_10'" json:"member_master_id_10"`
	MemberMasterId11 generic.Nullable[int32] `xorm:"json 'member_master_id_11'" json:"member_master_id_11"`
	MemberMasterId12 generic.Nullable[int32] `xorm:"json 'member_master_id_12'" json:"member_master_id_12"`
	SuitMasterId1    generic.Nullable[int32] `xorm:"json 'suit_master_id_1'" json:"suit_master_id_1"`
	SuitMasterId2    generic.Nullable[int32] `xorm:"json 'suit_master_id_2'" json:"suit_master_id_2"`
	SuitMasterId3    generic.Nullable[int32] `xorm:"json 'suit_master_id_3'" json:"suit_master_id_3"`
	SuitMasterId4    generic.Nullable[int32] `xorm:"json 'suit_master_id_4'" json:"suit_master_id_4"`
	SuitMasterId5    generic.Nullable[int32] `xorm:"json 'suit_master_id_5'" json:"suit_master_id_5"`
	SuitMasterId6    generic.Nullable[int32] `xorm:"json 'suit_master_id_6'" json:"suit_master_id_6"`
	SuitMasterId7    generic.Nullable[int32] `xorm:"json 'suit_master_id_7'" json:"suit_master_id_7"`
	SuitMasterId8    generic.Nullable[int32] `xorm:"json 'suit_master_id_8'" json:"suit_master_id_8"`
	SuitMasterId9    generic.Nullable[int32] `xorm:"json 'suit_master_id_9'" json:"suit_master_id_9"`
	SuitMasterId10   generic.Nullable[int32] `xorm:"json 'suit_master_id_10'" json:"suit_master_id_10"`
	SuitMasterId11   generic.Nullable[int32] `xorm:"json 'suit_master_id_11'" json:"suit_master_id_11"`
	SuitMasterId12   generic.Nullable[int32] `xorm:"json 'suit_master_id_12'" json:"suit_master_id_12"`
}

func (ulmd *UserLiveMvDeck) Id() int64 {
	return int64(ulmd.LiveMasterId)
}
