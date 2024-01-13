package client

import (
	"elichika/generic"
)

type BootstrapLiveCampaignInfo struct {
	LiveCampaignEndAt          generic.Nullable[int64] `json:"live_campaign_end_at"`
	LiveDailyCampaignEndAt     generic.Nullable[int64] `json:"live_daily_campaign_end_at"`
	LiveExtraCampaignEndAt     generic.Nullable[int64] `json:"live_extra_campaign_end_at"`
	LiveChallengeCampaignEndAt generic.Nullable[int64] `json:"live_challenge_campaign_end_at"`
}
