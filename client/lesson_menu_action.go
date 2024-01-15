package client

import (
	"elichika/generic"
)

type LessonMenuAction struct {
	CardMasterId                  int32                   `json:"card_master_id"`
	Position                      int32                   `json:"position"`
	IsAddedPassiveSkill           bool                    `json:"is_added_passive_skill"`
	IsAddedSpecialPassiveSkill    bool                    `json:"is_added_special_passive_skill"`
	IsRankupedPassiveSkill        bool                    `json:"is_rankuped_passive_skill"`
	IsRankupedSpecialPassiveSkill bool                    `json:"is_rankuped_special_passive_skill"`
	IsPromotedSkill               bool                    `json:"is_promoted_skill"`
	MaxRarity                     generic.Nullable[int32] `json:"max_rarity" enum:""`
	UpCount                       int32                   `json:"up_count"`
}
