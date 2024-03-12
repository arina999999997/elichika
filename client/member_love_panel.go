package client

import (
	"elichika/generic"

	"sort"
)

type MemberLovePanel struct {
	MemberId               int32                `xorm:"pk" json:"member_id"`
	MemberLovePanelCellIds generic.Array[int32] `xorm:"json" json:"member_love_panel_cell_ids"`
}

// TODO(love_panel): For now, member_love_panel is stored using user_id and member_id as pk, and the cells are the json
// it's possible to store using only cells of the last level or even level + bitset for reduced memory
func (mlp *MemberLovePanel) Fix() { // make the state correct, no matter what
	sort.Slice(mlp.MemberLovePanelCellIds.Slice, func(i, j int) bool {
		return mlp.MemberLovePanelCellIds.Slice[i] < mlp.MemberLovePanelCellIds.Slice[j]
	})
	for i := range mlp.MemberLovePanelCellIds.Slice {
		if i == 0 {
			continue
		}
		if mlp.MemberLovePanelCellIds.Slice[i] == mlp.MemberLovePanelCellIds.Slice[i-1] { // inconsistency, let's just reset it
			mlp.MemberLovePanelCellIds.Slice = []int32{}
			break
		}
	}
}
