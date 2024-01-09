package request

type SetFavoriteMemberRequest struct {
	MemberMasterId int `json:"member_master_id"`
}

type SetThemeRequest struct {
	MemberMasterId           int `json:"member_master_id"`
	SuitMasterId             int `json:"suit_master_id"`
	CustomBackgroundMasterId int `json:"custom_background_master_id"`
}

type FinishUserStoryMemberRequest struct {
	StoryMemberMasterId int32 `json:"story_member_master_id"`
	IsAutoMode          bool  `json:"is_auto_mode"`
}

type FinishUserStorySideRequest struct {
	StorySideMasterId int32 `json:"story_side_master_id"`
	IsAutoMode        bool  `json:"is_auto_mode"`
}
