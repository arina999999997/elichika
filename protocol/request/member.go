package request

type OpenMemberLovePanelRequest struct {
	MemberId               int     `json:"member_id"`
	MemberLovePanelId      int     `json:"member_love_panel_id"`
	MemberLovePanelCellIds []int32 `json:"member_love_panel_cell_ids"`
}
