package client

type UserInfoTriggerMemberLoveLevelUp struct {
	TriggerId       int64 `xorm:"pk 'trigger_id'" json:"trigger_id"`
	MemberMasterId  int32   `xorm:"'member_master_id'" json:"member_master_id"`
	BeforeLoveLevel int32   `json:"before_love_level"`
	// TODO(refactor): Remove IsNull
	IsNull          bool  `json:"-" xorm:"-"`
}

func (uitmllu *UserInfoTriggerMemberLoveLevelUp) Id() int64 {
	return uitmllu.TriggerId
}