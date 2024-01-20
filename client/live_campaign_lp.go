package client

type LiveCampaignLp struct {
	Id                 int32 `json:"id"`
	LiveDifficultyType int32 `json:"live_difficulty_type" enum:"LiveDifficultyType"`
	ConsumedLp         int32 `json:"consumed_lp"`
}
