package model

import (
	"elichika/generic"
)

type UserAccessory struct {
	UserAccessoryId    int64 `xorm:"pk 'user_accessory_id'" json:"user_accessory_id"`
	AccessoryMasterId  int   `xorm:"'accessory_master_id'" json:"accessory_master_id"`
	Level              int   `json:"level"`
	Exp                int   `json:"exp"`
	Grade              int   `json:"grade"`
	Attribute          int   `json:"attribute"`
	PassiveSkill1Id    int   `xorm:"'passive_skill_1_id'" json:"passive_skill_1_id"`
	PassiveSkill1Level int   `xorm:"'passive_skill_1_level'" json:"passive_skill_1_level"`
	PassiveSkill2Id    *int  `xorm:"'passive_skill_2_id'" json:"passive_skill_2_id"`
	PassiveSkill2Level int   `xorm:"'passive_skill_2_level'" json:"passive_skill_2_level"`
	IsLock             bool  `json:"is_lock"`
	IsNew              bool  `json:"is_new"`
	AcquiredAt         int64 `json:"acquired_at"` // unix second
	IsNull             bool  `json:"-" xorm:"-"`
}

func (ua *UserAccessory) Id() int64 {
	return ua.UserAccessoryId
}

func init() {
	TableNameToInterface["u_accessory"] = generic.UserIdWrapper[UserAccessory]{}
}
