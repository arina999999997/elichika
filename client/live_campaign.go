package client

import (
	"elichika/generic"
)

type LiveCampaign struct {
	LiveMasterId generic.Nullable[int32] `json:"live_master_id"`
	CampaignType int32                   `json:"campaign_type" enum:"LiveCampaignType"`
	ParamConst   generic.Nullable[int32] `json:"param_const"`
	ParamId      generic.Nullable[int32] `json:"param_id"`
	AppealText   LocalizedText           `json:"appeal_text"`
	EndAt        int64                   `json:"end_at"`
	DisplayOrder int32                   `json:"display_order"`
}
