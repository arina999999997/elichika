package client

import (
	"elichika/generic"
)

type LiveCampaignChangeDropContent struct {
	LiveId             int32                   `json:"live_id"`
	LiveDifficultyType int32                   `json:"live_difficulty_type"`
	DropContentGroupId int32                   `json:"drop_content_group_id"`
	Weekday            generic.Nullable[int32] `json:"weekday" enum:"weekday"`
}
