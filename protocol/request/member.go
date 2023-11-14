package request


type OpenMemberLovePanelRequest struct {
	MemberID               int   `json:"member_id"`
	MemberLovePanelID      int   `json:"member_love_panel_id"`
	MemberLovePanelCellIDs []int `json:"member_love_panel_cell_ids"`
}