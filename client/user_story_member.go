package client

type UserStoryMember struct {
	StoryMemberMasterId int32 `xorm:"pk 'story_member_master_id'" json:"story_member_master_id"`
	IsNew               bool  `xorm:"'is_new'" json:"is_new"`
	AcquiredAt          int64 `xorm:"'acquired_at'" json:"acquired_at"`
}
