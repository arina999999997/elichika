package model

// type LiveMVDeck struct {
// 	// unused, LiveMVDeck is not saved in server db
// 	UserID              int   `xorm:"pk 'user_id'" json:"-"`
// 	LiveMasterID        int   `xorm:"pk 'live_master_id'" json:"live_master_id"`
// 	LiveMvDeckType      int   `json:"live_mv_deck_type"` // 1 for original deck, 2 for custom deck
// 	MemberMasterIDByPos []int `xorm:"'member_master_id_by_pos'" json:"member_master_id_by_pos"`
// 	SuitMasterIDByPos   []int `xorm:"'suit_master_id_by_pos'" json:"suit_master_id_by_pos"`
// 	ViewStatusByPos     []int `json:"view_status_by_pos"`
// }

type UserLiveMvDeck struct {
	UserID           int  `xorm:"pk 'user_id'" json:"-"`
	LiveMasterID     int  `xorm:"pk 'live_master_id'" json:"live_master_id"`
	MemberMasterID1  *int `xorm:"'member_master_id_1'" json:"member_master_id_1"`
	MemberMasterID2  *int `xorm:"'member_master_id_2'" json:"member_master_id_2"`
	MemberMasterID3  *int `xorm:"'member_master_id_3'" json:"member_master_id_3"`
	MemberMasterID4  *int `xorm:"'member_master_id_4'" json:"member_master_id_4"`
	MemberMasterID5  *int `xorm:"'member_master_id_5'" json:"member_master_id_5"`
	MemberMasterID6  *int `xorm:"'member_master_id_6'" json:"member_master_id_6"`
	MemberMasterID7  *int `xorm:"'member_master_id_7'" json:"member_master_id_7"`
	MemberMasterID8  *int `xorm:"'member_master_id_8'" json:"member_master_id_8"`
	MemberMasterID9  *int `xorm:"'member_master_id_9'" json:"member_master_id_9"`
	MemberMasterID10 *int `xorm:"'member_master_id_10'" json:"member_master_id_10"`
	MemberMasterID11 *int `xorm:"'member_master_id_11'" json:"member_master_id_11"`
	MemberMasterID12 *int `xorm:"'member_master_id_12'" json:"member_master_id_12"`
	SuitMasterID1    *int `xorm:"'suit_master_id_1'" json:"suit_master_id_1"`
	SuitMasterID2    *int `xorm:"'suit_master_id_2'" json:"suit_master_id_2"`
	SuitMasterID3    *int `xorm:"'suit_master_id_3'" json:"suit_master_id_3"`
	SuitMasterID4    *int `xorm:"'suit_master_id_4'" json:"suit_master_id_4"`
	SuitMasterID5    *int `xorm:"'suit_master_id_5'" json:"suit_master_id_5"`
	SuitMasterID6    *int `xorm:"'suit_master_id_6'" json:"suit_master_id_6"`
	SuitMasterID7    *int `xorm:"'suit_master_id_7'" json:"suit_master_id_7"`
	SuitMasterID8    *int `xorm:"'suit_master_id_8'" json:"suit_master_id_8"`
	SuitMasterID9    *int `xorm:"'suit_master_id_9'" json:"suit_master_id_9"`
	SuitMasterID10   *int `xorm:"'suit_master_id_10'" json:"suit_master_id_10"`
	SuitMasterID11   *int `xorm:"'suit_master_id_11'" json:"suit_master_id_11"`
	SuitMasterID12   *int `xorm:"'suit_master_id_12'" json:"suit_master_id_12"`
}

func (ulmd *UserLiveMvDeck) ID() int64 {
	return int64(ulmd.LiveMasterID)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_live_mv_deck"] = UserLiveMvDeck{}
	TableNameToInterface["u_live_mv_deck_custom"] = UserLiveMvDeck{}
}
