package client

import (
	"elichika/generic"
)

type Caution struct {
	CautionId    int64                   `json:"caution_id"`
	CautionScene int32                   `json:"caution_scene" enum:"CautionSceneId"`
	CautionType  int32                   `json:"caution_type" enum:"CautionMethod"`
	StartAt      int64                   `json:"start_at"`
	EndAt        generic.Nullable[int64] `json:"end_at"`
	LookedAt     generic.Nullable[int64] `json:"looked_at"`
	Title        LocalizedText           `json:"title"`
	Message      LocalizedText           `json:"message"`
}
