package request

import (
	"elichika/generic"
)

type FinishEventStoryRequest struct {
	StoryEventMasterId int32                  `json:"story_event_master_id"`
	IsAutoMode         generic.Nullable[bool] `json:"is_auto_mode"`
}
