package client

import (
	"elichika/generic"
)

type LiveCardStat struct {
	GotVoltage             int32               `json:"got_voltage"`
	SkillTriggeredCount    int32               `json:"skill_triggered_count"`
	AppealCount            int32               `json:"appeal_count"`
	RecastSquadEffectCount int32               `json:"recast_squad_effect_count"`
	CardMasterId           int32               `json:"card_master_id"`
	BaseParameter          LiveCardParameter   `json:"base_parameter"`
	ActiveSkillValue       int32               `json:"active_skill_value"`
	PassiveSkillValues     generic.List[int32] `json:"passive_skill_values"`
}
