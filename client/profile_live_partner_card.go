package client

type ProfileLivePartnerCard struct {
	LivePartnerCategoryMasterId int32         `json:"live_partner_category_master_id"`
	PartnerCard                 OtherUserCard `json:"partner_card"`
}
