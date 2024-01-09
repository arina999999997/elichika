package client

type UserStoryEventHistory struct {
	StoryEventId int32 `xorm:"pk 'story_event_id'" json:"story_event_id"`
}

func (useh *UserStoryEventHistory) Id() int64 {
	return int64(useh.StoryEventId)
}
