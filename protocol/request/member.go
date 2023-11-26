package request

type OpenMemberLovePanelRequest struct {
	MemberID               int   `json:"member_id"`
	MemberLovePanelID      int   `json:"member_love_panel_id"`
	MemberLovePanelCellIDs []int `json:"member_love_panel_cell_ids"`
}

type UpdateUserCommunicationMemberDetailBadgeRequest struct {
	MemberMasterID                     int `json:"member_master_id"`
	CommunicationMemberDetailBadgeType int `json:"communication_member_detail_badge_type"`
}
