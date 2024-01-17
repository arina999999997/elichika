package client

type RankingUser struct {
	UserId                 int32         `json:"user_id"`
	Name                   LocalizedText `json:"name"`
	Rank                   int32         `json:"rank"`
	FavoriteCardMasterId   int32         `json:"favorite_card_master_id"`
	FavoriteCardLevel      int32         `json:"favorite_card_level"`
	IsAwakeningImage       bool          `json:"is_awakening_image"`
	IsAllTrainingActivated bool          `json:"is_all_training_activated"`
	EmblemId               int32         `json:"emblem_id"`
}
