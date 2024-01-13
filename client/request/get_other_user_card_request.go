package request

type GetOtherUserCardRequest struct {
	UserId       int32 `json:"user_id"`
	CardMasterId int32 `json:"card_master_id"`
}
