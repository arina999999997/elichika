package client

type LiveCampaignChangeDrop struct {
	Id                 int32 `json:"id"`
	LiveDifficultyType int32 `json:"live_difficulty_type" enum:"LiveDifficultyType"`
	DropContentGroupId int32 `json:"drop_content_group_id"`
}
