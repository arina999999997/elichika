package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchLiveMusicSelectResponse struct {
	WeekdayState               client.WeekdayState                          `json:"weekday_state"`
	LiveDailyList              generic.Array[client.LiveDaily]              `json:"live_daily_list"`
	LiveCampaignList           generic.Array[client.LiveCampaign]           `json:"live_campaign_list"`
	LiveDailyCampaignList      generic.Array[client.LiveDailyCampaign]      `json:"live_daily_campaign_list"`
	LiveCampaignLpList         generic.Array[client.LiveCampaignLp]         `json:"live_campaign_lp_list"`
	LiveCampaignChangeDropList generic.Array[client.LiveCampaignChangeDrop] `json:"live_campaign_change_drop_list"`
	LiveCampaignNotice         generic.Nullable[client.LiveCampaignNotice]  `json:"live_campaign_notice"`   // pointer
	LiveEventCommonInfo        generic.Nullable[client.LiveEventCommonInfo] `json:"live_event_common_info"` // pointer
	BannerInfo                 client.BootstrapBanner                       `json:"banner_info"`
	PickupInfo                 client.BootstrapPickupInfo                   `json:"pickup_info"`
	SteadyRankingInfo          generic.Nullable[client.SteadyRankingInfo]   `json:"steady_ranking_info"`
	UserModelDiff              *client.UserModel                            `json:"user_model_diff"`
}
