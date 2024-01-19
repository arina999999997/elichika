package request

type GradeUpCardRequest struct {
	CardMasterId int32 `json:"card_master_id"` // is actually named _CardMasterId
	ContentId    int32 `json:"content_id"`     // is actually named _ContentId
}
