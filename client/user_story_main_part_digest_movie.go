package client

type UserStoryMainPartDigestMovie struct {
	StoryMainPartMasterId int32 `xorm:"pk 'story_main_part_master_id'" json:"story_main_part_master_id"`
}

func (usmpdm *UserStoryMainPartDigestMovie) Id() int64 {
	return int64(usmpdm.StoryMainPartMasterId)
}
