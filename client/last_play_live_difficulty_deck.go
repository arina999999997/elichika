package client

import (
	"elichika/generic"
)

type LastPlayLiveDifficultyDeck struct {
	LiveDifficultyId int32                                                  `json:"live_difficulty_id"`
	Voltage          int32                                                  `json:"voltage"`
	IsCleared        bool                                                   `json:"is_cleared"`
	RecordedAt       int64                                                  `json:"recorded_at"`
	CardWithSuitDict generic.Dictionary[int32, int32]                       `xorm:"json" json:"card_with_suit_dict"`
	SquadDict        generic.Dictionary[int32, LastPlayLiveDifficultySquad] `xorm:"json" json:"squad_dict"`
}
