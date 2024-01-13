package client

import (
	"elichika/generic"
)

type MemberLovePanel struct {
	MemberId int32  `xorm:"pk" json:"member_id"`
	MemberLovePanelCellIds generic.Array[int32] `xorm:"json" json:"member_love_panel_cell_ids"`
}

// TODO(love_panel): For now, member_love_panel is stored using user_id and member_id as pk, and the cells are the json
// it's possible to store using only cells of the last level or even level + bitset for reduced memory