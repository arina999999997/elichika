package client

import (
	"elichika/generic"
)

type UserAccessory struct {
	UserAccessoryId    int64                   `xorm:"pk 'user_accessory_id'" json:"user_accessory_id"`
	AccessoryMasterId  int32                   `xorm:"'accessory_master_id'" json:"accessory_master_id"`
	Level              int32                   `json:"level"`
	Exp                int32                   `json:"exp"`
	Grade              int32                   `json:"grade"`
	Attribute          int32                   `json:"attribute" enum:"CardAttribute"`
	PassiveSkill1Id    generic.Nullable[int32] `xorm:"json 'passive_skill_1_id'" json:"passive_skill_1_id"`
	PassiveSkill1Level generic.Nullable[int32] `xorm:"json 'passive_skill_1_level'" json:"passive_skill_1_level"`
	PassiveSkill2Id    generic.Nullable[int32] `xorm:"json 'passive_skill_2_id'" json:"passive_skill_2_id"`
	PassiveSkill2Level generic.Nullable[int32] `xorm:"json 'passive_skill_2_level'" json:"passive_skill_2_level"`
	IsLock             bool                    `json:"is_lock"`
	IsNew              bool                    `json:"is_new"`
	AcquiredAt         int64                   `json:"acquired_at"` // unix second
}
