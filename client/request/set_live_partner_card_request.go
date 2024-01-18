package request

type SetLivePartnerCardRequest struct {
	LivePartnerCategoryId int32 `json:"live_partner_category_id"`
	CardMasterId          int32 `json:"card_master_id"`
}
