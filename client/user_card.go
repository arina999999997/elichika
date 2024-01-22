package client

type UserCard struct {
	CardMasterId               int32 `xorm:"pk 'card_master_id'" json:"card_master_id"`
	Level                      int32 `json:"level"`
	Exp                        int32 `json:"exp"`
	LovePoint                  int32 `json:"love_point"`
	IsFavorite                 bool  `json:"is_favorite"`
	IsAwakening                bool  `json:"is_awakening"`
	IsAwakeningImage           bool  `json:"is_awakening_image"`
	IsAllTrainingActivated     bool  `json:"is_all_training_activated"`
	TrainingActivatedCellCount int32 `json:"training_activated_cell_count"`
	MaxFreePassiveSkill        int32 `json:"max_free_passive_skill"`
	Grade                      int32 `json:"grade"`
	TrainingLife               int32 `json:"training_life"`
	TrainingAttack             int32 `json:"training_attack"`
	TrainingDexterity          int32 `json:"training_dexterity"`
	ActiveSkillLevel           int32 `json:"active_skill_level"`
	PassiveSkillALevel         int32 `json:"passive_skill_a_level"`
	PassiveSkillBLevel         int32 `json:"passive_skill_b_level"`
	PassiveSkillCLevel         int32 `json:"passive_skill_c_level"`
	AdditionalPassiveSkill1Id  int32 `xorm:"'additional_passive_skill_1_id'" json:"additional_passive_skill_1_id"`
	AdditionalPassiveSkill2Id  int32 `xorm:"'additional_passive_skill_2_id'" json:"additional_passive_skill_2_id"`
	AdditionalPassiveSkill3Id  int32 `xorm:"'additional_passive_skill_3_id'" json:"additional_passive_skill_3_id"`
	AdditionalPassiveSkill4Id  int32 `xorm:"'additional_passive_skill_4_id'" json:"additional_passive_skill_4_id"`
	AcquiredAt                 int32 `json:"acquired_at"` // should be int64 but client use int32
	IsNew                      bool  `json:"is_new"`
}

func (uc *UserCard) Id() int64 {
	return int64(uc.CardMasterId)
}
