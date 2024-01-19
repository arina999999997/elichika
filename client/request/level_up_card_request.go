package request

type LevelUpCardRequest struct {
	AdditionalLevel int32 `json:"additional_level"`
	CardMasterId    int32 `json:"card_master_id"`
}
