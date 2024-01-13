package client

import (
	"elichika/generic"
)

type OtherUserCard struct {
	CardMasterId int32 `json:"card_master_id"`
	Level int32  `json:"level"`
	Grade int32  `json:"grade"`
	LoveLevel int32 `json:"love_level"`
	IsAwakening bool `json:"is_awakening"`
	IsAwakeningImage bool `json:"is_awakening_image"`
	IsAllTrainingActivated bool `json:"is_all_training_activated"`
	ActiveSkillLevel int32 `json:"active_skill_level"`
	PassiveSkillLevels generic.Array[int32] `json:"passive_skill_levels"`
	AdditionalPassiveSkillIds generic.Array[int32] `json:"additional_passive_skill_ids"`
	MaxFreePassiveSkill int32 `json:"max_free_passive_skill"`
	TrainingStamina int32 `json:"training_stamina"`
	TrainingAppeal int32 `json:"training_appeal"`
	TrainingTechnique int32 `json:"training_technique"`
	MemberLovePanels generic.Array[MemberLovePanel] `json:"member_love_panels"`
}