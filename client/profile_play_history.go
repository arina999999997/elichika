package client

import (
	"elichika/generic"
)

type ProfilePlayHistory struct {
	LivePlayCount          generic.Dictionary[int32, int32] `json:"live_play_count" enum:"LiveDifficultyType"`
	LiveClearCount         generic.Dictionary[int32, int32] `json:"live_clear_count" enum:"LiveDifficultyType"`
	JoinedLiveCardRanking  generic.Array[ProfileUserCard]   `json:"joined_live_card_ranking"`
	PlaySkillCardRanking   generic.Array[ProfileUserCard]   `json:"play_skill_card_ranking"`
	MaxScoreLiveDifficulty ProfileLiveDifficulty            `json:"max_score_live_difficulty"`
	MaxComboLiveDifficulty ProfileLiveDifficulty            `json:"max_combo_live_difficulty"`
}
