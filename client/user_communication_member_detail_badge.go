package client

type UserCommunicationMemberDetailBadge struct {
	MemberMasterId     int32 `xorm:"pk 'member_master_id'" json:"member_master_id"`
	IsStoryMemberBadge bool  `xorm:"'is_story_member_badge'" json:"is_story_member_badge"`
	IsStorySideBadge   bool  `xorm:"'is_story_side_badge'" json:"is_story_side_badge"`
	IsVoiceBadge       bool  `xorm:"'is_voice_badge'" json:"is_voice_badge"`
	IsThemeBadge       bool  `xorm:"'is_theme_badge'" json:"is_theme_badge"`
	IsCardBadge        bool  `xorm:"'is_card_badge'" json:"is_card_badge"`
	IsMusicBadge       bool  `xorm:"'is_music_badge'" json:"is_music_badge"`
}

func (ucmdb *UserCommunicationMemberDetailBadge) Id() int64 {
	return int64(ucmdb.MemberMasterId)
}
