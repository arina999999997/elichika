package model

import (
	"elichika/client"
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

func init() {

	type DbLiveRecord struct {
		client.UserLiveDifficulty `xorm:"extends"`
		Voltage                   int   `xorm:"'last_clear_voltage'" json:"voltage"`
		IsCleared                 bool  `xorm:"'last_clear_is_cleared'" json:"is_cleared"`
		RecordedAt                int64 `xorm:"'last_clear_recorded_at'" json:"recorded_at"`
		CardWithSuitDict          []int `xorm:"'last_clear_cards_and_suits'" json:"card_with_suit_dict"`
		SquadDict                 []any `xorm:"'squad_dict'" json:"squad_dict"`
	}
	TableNameToInterface["u_live_record"] = generic.UserIdWrapper[DbLiveRecord]{}
}
