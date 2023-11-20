package model

type LastPlayLiveDifficultyDeck struct {
	UserID           int   `xorm:"pk 'user_id'" json:"-"`
	LiveDifficultyID int   `xorm:"pk 'live_difficulty_id'" json:"live_difficulty_id"`
	Voltage          int   `xorm:"'last_clear_voltage'" json:"voltage"`
	IsCleared        bool  `xorm:"'last_clear_is_cleared'" json:"is_cleared"`
	RecordedAt       int64 `xorm:"'last_clear_recorded_at'" json:"recorded_at"`
	CardWithSuitDict []int `xorm:"'last_clear_cards_and_suits'" json:"card_with_suit_dict"`
	SquadDict        []any `xorm:"'squad_dict'" json:"squad_dict"`
}

type UserLiveDifficulty struct {
	UserID                        int  `xorm:"pk 'user_id'" json:"-"`
	LiveDifficultyID              int  `xorm:"pk 'live_difficulty_id'" json:"live_difficulty_id"`
	MaxScore                      int  `xorm:"'max_score'" json:"max_score"`
	MaxCombo                      int  `xorm:"'max_combo'" json:"max_combo"`
	PlayCount                     int  `xorm:"'play_count'" json:"play_count"`   // live start count
	ClearCount                    int  `xorm:"'clear_count'" json:"clear_count"` // live finish and cleared
	CancelCount                   int  `xorm:"-" json:"cancel_count"`            // unused, always 0
	NotClearedCount               int  `xorm:"-" json:"not_cleared_count"`
	IsFullCombo                   bool `xorm:"-" json:"is_full_combo"`                                                     // isn't used, at least not when autoplay is used
	ClearedDifficultyAchievement1 *int `xorm:"'cleared_difficulty_achievement_1'" json:"cleared_difficulty_achievement_1"` // 1 if cleared, null if not?
	ClearedDifficultyAchievement2 *int `xorm:"'cleared_difficulty_achievement_2'" json:"cleared_difficulty_achievement_2"` // 1 if cleared, null if not?
	ClearedDifficultyAchievement3 *int `xorm:"'cleared_difficulty_achievement_3'" json:"cleared_difficulty_achievement_3"` // 1 if cleared, null if not?
	EnableAutoplay                bool `xorm:"'enable_autoplay'" json:"enable_autoplay"`                                   // can autoplay?
	IsAutoplay                    bool `xorm:"'is_autoplay'" json:"is_autoplay"`                                           // is using autoplay?
	IsNew                         bool `xorm:"'is_new'" json:"is_new"`
}

func (uld *UserLiveDifficulty) ID() int64 {
	return int64(uld.LiveDifficultyID)
}


func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	type DbLiveRecord struct {
		UserLiveDifficulty `xorm:"extends"`
		Voltage                  int   `xorm:"'last_clear_voltage'" json:"voltage"`
		IsCleared                bool  `xorm:"'last_clear_is_cleared'" json:"is_cleared"`
		RecordedAt               int64 `xorm:"'last_clear_recorded_at'" json:"recorded_at"`
		CardWithSuitDict         []int `xorm:"'last_clear_cards_and_suits'" json:"card_with_suit_dict"`
		SquadDict                []any `xorm:"'squad_dict'" json:"squad_dict"`
	}
	TableNameToInterface["u_live_record"]=DbLiveRecord{}
}