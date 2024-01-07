package response

import (
	"elichika/model"
)

type RankingUser struct {
	UserId                 int                 `json:"user_id"`
	UserName               model.LocalizedText `json:"user_name"`
	UserRank               int                 `json:"user_rank"`
	CardMasterId           int                 `json:"card_master_id"`
	Level                  int                 `json:"level"`
	IsAwakening            bool                `json:"is_awakening"`
	IsAllTrainingActivated bool                `json:"is_all_training_activated"`
	EmblemMasterId         int                 `json:"emblem_master_id"`
}
