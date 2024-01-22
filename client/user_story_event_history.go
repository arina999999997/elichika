package client

type UserStoryEventHistory struct {
	StoryEventId int32 `xorm:"pk 'story_event_id'" json:"story_event_id"`
}
