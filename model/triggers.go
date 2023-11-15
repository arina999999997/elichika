package model

type TriggerBasic struct {
	UserID          int     `xorm:"pk 'user_id'" json:"-"`
	TriggerID       int64   `xorm:"pk 'trigger_id'" json:"trigger_id"` // use nano timestamp
	InfoTriggerType int     `json:"info_trigger_type"`
	LimitAt         *int64  `json:"limit_at"` // seems like some sort of timed timestamp, probably for event popup
	Description     *string `json:"description"`
	ParamInt        int     `json:"param_int"`
	IsNull          bool    `json:"-" xorm:"-"`
}

func (obj *TriggerBasic) ID() int64 {
	return obj.TriggerID
}

type TriggerCardGradeUp struct {
	UserID               int   `xorm:"pk 'user_id'" json:"-"`
	TriggerID            int64 `xorm:"pk 'trigger_id'" json:"trigger_id"` // use nano timestamp
	CardMasterID         int   `xorm:"pk 'card_master_id'" json:"card_master_id"`
	BeforeLoveLevelLimit int   `json:"before_love_level_limit"`
	AfterLoveLevelLimit  int   `json:"after_love_level_limit"`
	IsNull               bool  `json:"-" xorm:"-"`
}

func (obj *TriggerCardGradeUp) ID() int64 {
	return obj.TriggerID
}

type TriggerMemberLoveLevelUp struct {
	// special thanks to https://github.com/Francesco149/todokete/blob/master/Todokete.kt
	UserID          int   `xorm:"pk 'user_id'" json:"-"`
	TriggerID       int64 `xorm:"pk 'trigger_id'" json:"trigger_id"` // use nano timestamp
	MemberMasterID  int   `xorm:"'member_master_id'" json:"member_master_id"`
	BeforeLoveLevel int   `json:"before_love_level"`
	IsNull          bool  `json:"-" xorm:"-"`
}

func (obj *TriggerMemberLoveLevelUp) ID() int64 {
	return obj.TriggerID
}

type TriggerReadReq struct {
	TriggerID int64 `json:"trigger_id"` // same for all trigger, for now
}
