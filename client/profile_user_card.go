package client

type ProfileUserCard struct {
	CardMasterId           int32 `json:"card_master_id"`
	Level                  int32 `json:"level"`
	IsAwakeningImage       bool  `json:"is_awakening_image"`
	IsAllTrainingActivated bool  `json:"is_all_training_activated"`
	LiveJoinCount          int32 `json:"live_join_count"`
	ActiveSkillPlayCount   int32 `json:"active_skill_play_count"`
}
