package client

import (
	"elichika/generic"
)

type BootstrapPickupInfo struct {
	ActiveEvent      generic.Nullable[BootstrapPickupEventInfo] `json:"active_event"`
	LiveCampaignInfo BootstrapLiveCampaignInfo                  `json:"live_campaign_info"`
	IsLessonCampaign bool                                       `json:"is_lesson_campaign"`
	AppealGachas     generic.Array[TextureStruktur]             `json:"appeal_gachas"`
	IsShopSale       bool                                       `json:"is_shop_sale"`
	IsSnsCoinSale    bool                                       `json:"is_sns_coin_sale"`
}
