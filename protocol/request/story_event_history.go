package request

type UnlockStoryRequest struct {
	EventStoryMasterID int `json:"event_story_master_id"`
}

type FinishStoryRequest struct {
	EventStoryMasterID int  `json:"event_story_master_id"`
	IsAutoMode         bool `json:"is_auto_mode"`
}
