package model

import (
	"elichika/generic"
)

type TriggerBasic struct {
	TriggerId       int64   `xorm:"pk 'trigger_id'" json:"trigger_id"` // use nano timestamp
	InfoTriggerType int     `json:"info_trigger_type"`
	LimitAt         *int64  `json:"limit_at"` // seems like some sort of timed timestamp, probably for event popup
	Description     *string `json:"description"`
	ParamInt        int     `json:"param_int"`
	IsNull          bool    `json:"-" xorm:"-"`
}

func (obj *TriggerBasic) Id() int64 {
	return obj.TriggerId
}

type TriggerCardGradeUp struct {
	TriggerId            int64 `xorm:"pk 'trigger_id'" json:"trigger_id"` // use nano timestamp
	CardMasterId         int   `xorm:"pk 'card_master_id'" json:"card_master_id"`
	BeforeLoveLevelLimit int   `json:"before_love_level_limit"`
	AfterLoveLevelLimit  int   `json:"after_love_level_limit"`
	IsNull               bool  `json:"-" xorm:"-"`
}

func (obj *TriggerCardGradeUp) Id() int64 {
	return obj.TriggerId
}

type TriggerMemberLoveLevelUp struct {
	// special thanks to https://github.com/Francesco149/todokete/blob/master/Todokete.kt
	TriggerId       int64 `xorm:"pk 'trigger_id'" json:"trigger_id"` // use nano timestamp
	MemberMasterId  int   `xorm:"'member_master_id'" json:"member_master_id"`
	BeforeLoveLevel int   `json:"before_love_level"`
	IsNull          bool  `json:"-" xorm:"-"`
}

func (obj *TriggerMemberLoveLevelUp) Id() int64 {
	return obj.TriggerId
}

type TriggerMemberGuildSupportItemExpired struct {
	// always present even if it's not actually expire
	TriggerId int64 `xorm:"pk 'trigger_id'" json:"trigger_id"` // use nano timestamp
	ResetAt   int   `xorm:"'reset_at'" json:"reset_at"`        // use unit timestamp
	IsNull    bool  `json:"-" xorm:"-"`
}

func (obj *TriggerMemberGuildSupportItemExpired) Id() int64 {
	return obj.TriggerId
}

type TriggerReadReq struct {
	TriggerId int64 `json:"trigger_id"` // same for all trigger, for now
}

func init() {

	TableNameToInterface["u_trigger_basic"] = generic.UserIdWrapper[TriggerBasic]{}
	TableNameToInterface["u_trigger_card_grade_up"] = generic.UserIdWrapper[TriggerCardGradeUp]{}
	TableNameToInterface["u_trigger_member_love_level_up"] = generic.UserIdWrapper[TriggerMemberLoveLevelUp]{}
	TableNameToInterface["u_trigger_member_guild_support_item_expired"] = generic.UserIdWrapper[TriggerMemberGuildSupportItemExpired]{}
}
