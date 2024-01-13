package request

type SetThemeRequest struct {
	MemberMasterId           int32 `json:"member_master_id"`
	SuitMasterId             int32 `json:"suit_master_id"`
	CustomBackgroundMasterId int32 `json:"custom_background_master_id"`
}
