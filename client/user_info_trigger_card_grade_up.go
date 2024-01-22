package client

type UserInfoTriggerCardGradeUp struct {
	TriggerId            int64 `xorm:"pk 'trigger_id'" json:"trigger_id"` // use nano timestamp
	CardMasterId         int32 `xorm:"'card_master_id'" json:"card_master_id"`
	BeforeLoveLevelLimit int32 `json:"before_love_level_limit"`
	AfterLoveLevelLimit  int32 `json:"after_love_level_limit"`
}
