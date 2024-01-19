package request

import (
	"elichika/generic"
)

type FinishStoryEventHistoryRequest struct {
	EventStoryMasterId int32                  `json:"event_story_master_id"`
	IsAutoMode         generic.Nullable[bool] `json:"is_auto_mode"`
}
