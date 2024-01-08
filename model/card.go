package model

// CardAwakeningReq ...
type CardAwakeningReq struct {
	CardMasterId     int  `json:"card_master_id"`
	IsAwakeningImage bool `json:"is_awakening_image"`
}

// CardFavoriteReq ...
type CardFavoriteReq struct {
	CardMasterId int  `json:"card_master_id"`
	IsFavorite   bool `json:"is_favorite"`
}

// UserCardReq ...
type UserCardReq struct {
	UserId       int64 `json:"user_id"`
	CardMasterId int64 `json:"card_master_id"`
}

// PartnerCard (Other user's card)
type PartnerCardInfo struct {
	CardMasterId              int   `xorm:"'card_master_id' default 0" json:"card_master_id"`
	Level                     int   `json:"level"`
	Grade                     int   `json:"grade"`
	LoveLevel                 int   `json:"love_level"`
	IsAwakening               bool  `json:"is_awakening"`
	IsAwakeningImage          bool  `json:"is_awakening_image"`
	IsAllTrainingActivated    bool  `json:"is_all_training_activated"`
	ActiveSkillLevel          int   `json:"active_skill_level"`
	PassiveSkillLevels        []int `json:"passive_skill_levels"`
	AdditionalPassiveSkillIds []int `json:"additional_passive_skill_ids"`
	MaxFreePassiveSkill       int   `json:"max_free_passive_skill"`
	TrainingStamina           int   `json:"training_stamina"`
	TrainingAppeal            int   `json:"training_appeal"`
	TrainingTechnique         int   `json:"training_technique"`
	MemberLovePanels          []int `json:"member_love_panels"`
}

type CardPlayInfo struct {
	CardMasterId           int  `xorm:"'card_master_id'" json:"card_master_id"`
	Level                  int  `json:"level"`
	IsAwakeningImage       bool `json:"is_awakening_image"`
	IsAllTrainingActivated bool `json:"is_all_training_activated"`
	LiveJoinCount          int  `xorm:"'live_join_count' default 0" json:"live_join_count"`
	ActiveSkillPlayCount   int  `xorm:"'active_skill_play_count' default 0" json:"active_skill_play_count"`
}
