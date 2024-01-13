package request

type ChangeFavoriteRequest struct {
	CardMasterId int32 `json:"card_master_id"`
	IsFavorite   bool  `json:"is_favorite"`
}
