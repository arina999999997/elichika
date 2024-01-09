package request

type UnlockStoryRequest struct {
	EventStoryMasterId int32 `json:"event_story_master_id"`
}

type FinishStoryRequest struct {
	EventStoryMasterId int32 `json:"event_story_master_id"`
	IsAutoMode         bool  `json:"is_auto_mode"`
}
