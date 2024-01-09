package client

type UserStoryMain struct {
	StoryMainMasterId int32 `xorm:"pk 'story_main_master_id'" json:"story_main_master_id"`
}

func (usm *UserStoryMain) Id() int64 {
	return int64(usm.StoryMainMasterId)
}
