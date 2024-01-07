package model

import (
	"elichika/generic"
)

// type LiveMVDeck struct {
// 	// unused, LiveMVDeck is not saved in server db
// 	LiveMasterId        int   `xorm:"pk 'live_master_id'" json:"live_master_id"`
// 	LiveMvDeckType      int   `json:"live_mv_deck_type"` // 1 for original deck, 2 for custom deck
// 	MemberMasterIdByPos []int `xorm:"'member_master_id_by_pos'" json:"member_master_id_by_pos"`
// 	SuitMasterIdByPos   []int `xorm:"'suit_master_id_by_pos'" json:"suit_master_id_by_pos"`
// 	ViewStatusByPos     []int `json:"view_status_by_pos"`
// }

type UserLiveMvDeck struct {
	LiveMasterId     int  `xorm:"pk 'live_master_id'" json:"live_master_id"`
	MemberMasterId1  *int `xorm:"'member_master_id_1'" json:"member_master_id_1"`
	MemberMasterId2  *int `xorm:"'member_master_id_2'" json:"member_master_id_2"`
	MemberMasterId3  *int `xorm:"'member_master_id_3'" json:"member_master_id_3"`
	MemberMasterId4  *int `xorm:"'member_master_id_4'" json:"member_master_id_4"`
	MemberMasterId5  *int `xorm:"'member_master_id_5'" json:"member_master_id_5"`
	MemberMasterId6  *int `xorm:"'member_master_id_6'" json:"member_master_id_6"`
	MemberMasterId7  *int `xorm:"'member_master_id_7'" json:"member_master_id_7"`
	MemberMasterId8  *int `xorm:"'member_master_id_8'" json:"member_master_id_8"`
	MemberMasterId9  *int `xorm:"'member_master_id_9'" json:"member_master_id_9"`
	MemberMasterId10 *int `xorm:"'member_master_id_10'" json:"member_master_id_10"`
	MemberMasterId11 *int `xorm:"'member_master_id_11'" json:"member_master_id_11"`
	MemberMasterId12 *int `xorm:"'member_master_id_12'" json:"member_master_id_12"`
	SuitMasterId1    *int `xorm:"'suit_master_id_1'" json:"suit_master_id_1"`
	SuitMasterId2    *int `xorm:"'suit_master_id_2'" json:"suit_master_id_2"`
	SuitMasterId3    *int `xorm:"'suit_master_id_3'" json:"suit_master_id_3"`
	SuitMasterId4    *int `xorm:"'suit_master_id_4'" json:"suit_master_id_4"`
	SuitMasterId5    *int `xorm:"'suit_master_id_5'" json:"suit_master_id_5"`
	SuitMasterId6    *int `xorm:"'suit_master_id_6'" json:"suit_master_id_6"`
	SuitMasterId7    *int `xorm:"'suit_master_id_7'" json:"suit_master_id_7"`
	SuitMasterId8    *int `xorm:"'suit_master_id_8'" json:"suit_master_id_8"`
	SuitMasterId9    *int `xorm:"'suit_master_id_9'" json:"suit_master_id_9"`
	SuitMasterId10   *int `xorm:"'suit_master_id_10'" json:"suit_master_id_10"`
	SuitMasterId11   *int `xorm:"'suit_master_id_11'" json:"suit_master_id_11"`
	SuitMasterId12   *int `xorm:"'suit_master_id_12'" json:"suit_master_id_12"`
}

func (ulmd *UserLiveMvDeck) Id() int64 {
	return int64(ulmd.LiveMasterId)
}

func init() {

	TableNameToInterface["u_live_mv_deck"] = generic.UserIdWrapper[UserLiveMvDeck]{}
	TableNameToInterface["u_live_mv_deck_custom"] = generic.UserIdWrapper[UserLiveMvDeck]{}
}
