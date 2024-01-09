package model

import (
	"elichika/client"
)

type RewardDrop struct { // unused
	DropColor int32          `json:"drop_color"`
	Content   client.Content `json:"content"`
	IsRare    bool           `json:"is_rare"`
	BonusType *int           `json:"bonus_type"`
}
