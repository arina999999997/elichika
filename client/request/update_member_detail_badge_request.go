package request

type UpdateMemberDetailBadgeRequest struct {
	MemberMasterId                     int32 `json:"member_master_id"`
	CommunicationMemberDetailBadgeType int32 `json:"communication_member_detail_badge_type" enum:"CommunicationMemberDetailBadgeType"`
}
