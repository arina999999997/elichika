package model

import (
	"elichika/generic"
)

type LastPlayLiveDifficultyDeck struct {
	LiveDifficultyId int   `xorm:"pk 'live_difficulty_id'" json:"live_difficulty_id"`
	Voltage          int   `xorm:"'last_clear_voltage'" json:"voltage"`
	IsCleared        bool  `xorm:"'last_clear_is_cleared'" json:"is_cleared"`
	RecordedAt       int64 `xorm:"'last_clear_recorded_at'" json:"recorded_at"`
	CardWithSuitDict []int `xorm:"'last_clear_cards_and_suits'" json:"card_with_suit_dict"`
	SquadDict        []any `xorm:"'squad_dict'" json:"squad_dict"`
}

type UserLiveDifficulty struct {
	LiveDifficultyId              int32 `xorm:"pk 'live_difficulty_id'" json:"live_difficulty_id"`
	MaxScore                      int32 `xorm:"'max_score'" json:"max_score"`
	MaxCombo                      int32 `xorm:"'max_combo'" json:"max_combo"`
	PlayCount                     int32 `xorm:"'play_count'" json:"play_count"`   // live start count
	ClearCount                    int32 `xorm:"'clear_count'" json:"clear_count"` // live finish and cleared
	CancelCount                   int32 `xorm:"-" json:"cancel_count"`            // unused, always 0
	NotClearedCount               int32 `xorm:"-" json:"not_cleared_count"`
	IsFullCombo                   bool  `xorm:"-" json:"is_full_combo"`                                                     // isn't used, at least not when autoplay is used
	ClearedDifficultyAchievement1 *int  `xorm:"'cleared_difficulty_achievement_1'" json:"cleared_difficulty_achievement_1"` // 1 if cleared, null if not?
	ClearedDifficultyAchievement2 *int  `xorm:"'cleared_difficulty_achievement_2'" json:"cleared_difficulty_achievement_2"` // 1 if cleared, null if not?
	ClearedDifficultyAchievement3 *int  `xorm:"'cleared_difficulty_achievement_3'" json:"cleared_difficulty_achievement_3"` // 1 if cleared, null if not?
	EnableAutoplay                bool  `xorm:"'enable_autoplay'" json:"enable_autoplay"`                                   // can autoplay?
	IsAutoplay                    bool  `xorm:"'is_autoplay'" json:"is_autoplay"`                                           // is using autoplay?
	IsNew                         bool  `xorm:"'is_new'" json:"is_new"`
}

func (uld *UserLiveDifficulty) Id() int64 {
	return int64(uld.LiveDifficultyId)
}

func init() {

	type DbLiveRecord struct {
		UserLiveDifficulty `xorm:"extends"`
		Voltage            int   `xorm:"'last_clear_voltage'" json:"voltage"`
		IsCleared          bool  `xorm:"'last_clear_is_cleared'" json:"is_cleared"`
		RecordedAt         int64 `xorm:"'last_clear_recorded_at'" json:"recorded_at"`
		CardWithSuitDict   []int `xorm:"'last_clear_cards_and_suits'" json:"card_with_suit_dict"`
		SquadDict          []any `xorm:"'squad_dict'" json:"squad_dict"`
	}
	TableNameToInterface["u_live_record"] = generic.UserIdWrapper[DbLiveRecord]{}
}
