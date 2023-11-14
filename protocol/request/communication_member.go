package request

type SetFavoriteMemberRequest struct {
	MemberMasterID int `json:"member_master_id"`
}

type SetThemeRequest struct {
	MemberMasterID           int `json:"member_master_id"`
	SuitMasterID             int `json:"suit_master_id"`
	CustomBackgroundMasterID int `json:"custom_background_master_id"`
}
