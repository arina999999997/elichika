package request

type SetFavoriteMemberRequest struct {
	MemberMasterID int `json:"member_master_id"`
}

type SetThemeRequest struct {
	MemberMasterID           int `json:"member_master_id"`
	SuitMasterID             int `json:"suit_master_id"`
	CustomBackgroundMasterID int `json:"custom_background_master_id"`
}

type FinishUserStoryMemberRequest struct {
	StoryMemberMasterID           int `json:"story_member_master_id"`
	IsAutoMode bool `json:"is_auto_mode"`
}

type FinishUserStorySideRequest struct {
	StorySideMasterID           int `json:"story_side_master_id"`
	IsAutoMode bool `json:"is_auto_mode"`
}