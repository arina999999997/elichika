package client

import (
	"elichika/generic"
)

type LiveStageGimmick struct {
	GimmickMasterId    int32                   `json:"gimmick_master_id"`
	ConditionMasterId1 int32                   `json:"condition_master_id_1"`
	ConditionMasterId2 generic.Nullable[int32] `json:"condition_master_id_2"`
	SkillMasterId      int32                   `json:"skill_master_id"`
	UniqId             int32                   `json:"uniq_id"`
}
