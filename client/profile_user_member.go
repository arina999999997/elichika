package client

type ProfileUserMember struct {
	MemberMasterId                int32 `json:"member_master_id"`
	LoveLevel                     int32 `json:"love_level"`
	LovePointLimit                int32 `json:"love_point_limit"`
	OwnedCardCount                int32 `json:"owned_card_count"`
	AllTrainingActivatedCardCount int32 `json:"all_training_activated_card_count"`
}
