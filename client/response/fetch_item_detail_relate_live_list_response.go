package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchItemDetailRelateLiveListResponse struct {
	WeekdayState                  client.WeekdayState                                 `json:"weekday_state"`
	LiveCampaignDropContents      generic.Array[client.LiveCampaignChangeDropContent] `json:"live_campaign_drop_contents"`
	LiveDailyCampaignDropContents generic.Array[client.LiveCampaignChangeDropContent] `json:"live_daily_campaign_drop_contents"`
}
