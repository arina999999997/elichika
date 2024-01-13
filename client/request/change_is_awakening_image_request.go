package request

type ChangeIsAwakeningImageRequest struct {
	CardMasterId     int32 `json:"card_master_id"`
	IsAwakeningImage bool  `json:"is_awakening_image"`
}
