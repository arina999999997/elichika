package client

import (
	"elichika/generic"
)

type OtherUserAccessory struct {
	UserAccessoryId    int64                   `json:"user_accessory_id"`
	AccessoryMasterId  int32                   `json:"accessory_master_id"`
	Level              int32                   `json:"level"`
	Grade              int32                   `json:"grade"`
	Attribute          int32                   `json:"attribute"`
	PassiveSkill1Id    generic.Nullable[int32] `xorm:"json 'passive_skill_1_id'" json:"passive_skill_1_id"`
	PassiveSkill1Level generic.Nullable[int32] `xorm:"json 'passive_skill_1_level'" json:"passive_skill_1_level"`
	PassiveSkill2Id    generic.Nullable[int32] `xorm:"json 'passive_skill_2_id'" json:"passive_skill_2_id"`
	PassiveSkill2Level generic.Nullable[int32] `xorm:"json 'passive_skill_2_level'" json:"passive_skill_2_level"`
}

func (ua UserAccessory) ToOtherUserAccessory() OtherUserAccessory {
	return OtherUserAccessory{
		UserAccessoryId:    ua.UserAccessoryId,
		AccessoryMasterId:  ua.AccessoryMasterId,
		Level:              ua.Level,
		Grade:              ua.Grade,
		Attribute:          ua.Attribute,
		PassiveSkill1Id:    ua.PassiveSkill1Id,
		PassiveSkill1Level: ua.PassiveSkill1Level,
		PassiveSkill2Id:    ua.PassiveSkill2Id,
		PassiveSkill2Level: ua.PassiveSkill2Level,
	}
}

func (oua *OtherUserAccessory) FromUserAccessory(ua UserAccessory) {
	oua.UserAccessoryId = ua.UserAccessoryId
	oua.AccessoryMasterId = ua.AccessoryMasterId
	oua.Level = ua.Level
	oua.Grade = ua.Grade
	oua.Attribute = ua.Attribute
	oua.PassiveSkill1Id = ua.PassiveSkill1Id
	oua.PassiveSkill1Level = ua.PassiveSkill1Level
	oua.PassiveSkill2Id = ua.PassiveSkill2Id
	oua.PassiveSkill2Level = ua.PassiveSkill2Level
}
