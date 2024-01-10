package request

type OpenMemberLovePanelRequest struct {
	MemberId               int     `json:"member_id"`
	MemberLovePanelId      int     `json:"member_love_panel_id"`
	MemberLovePanelCellIds []int32 `json:"member_love_panel_cell_ids"`
}

type UpdateUserCommunicationMemberDetailBadgeRequest struct {
	MemberMasterId                     int32 `json:"member_master_id"`
	CommunicationMemberDetailBadgeType int   `json:"communication_member_detail_badge_type"`
}
