package client

// TODO(event): not actually implemented, just for archival purpose
type UserEventMarathon struct { // Story ranking event
	EventMasterId     int32 `xorm:"pk 'event_master_id'" json:"event_master_id"`
	EventPoint        int32 `xorm:"'event_point'" json:"event_point"`
	OpenedStoryNumber int32 `xorm:"'opened_story_number'" json:"opened_story_number"`
	ReadStoryNumber   int32 `xorm:"'read_story_number'" json:"read_story_number"`
}
