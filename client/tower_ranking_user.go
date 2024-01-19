package client

type TowerRankingUser struct {
	UserId                 int32         `json:"user_id"`
	UserName               LocalizedText `json:"user_name"`
	UserRank               int32         `json:"user_rank"`
	CardMasterId           int32         `json:"card_master_id"`
	Level                  int32         `json:"level"`
	IsAwakening            bool          `json:"is_awakening"`
	IsAllTrainingActivated bool          `json:"is_all_training_activated"`
	EmblemMasterId         int32         `json:"emblem_master_id"`
}
