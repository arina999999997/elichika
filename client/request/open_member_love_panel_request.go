package request

import (
	"elichika/generic"
)

type OpenMemberLovePanelRequest struct {
	MemberId               int32                `json:"member_id"`
	MemberLovePanelId      int32                `json:"member_love_panel_id"`
	MemberLovePanelCellIds generic.Array[int32] `json:"member_love_panel_cell_ids"`
}
