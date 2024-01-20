package client

import (
	"elichika/generic"
)

type LiveDropContent struct {
	DropColor int32                   `json:"drop_color"`
	Content   Content                 `json:"content"`
	IsRare    bool                    `json:"is_rare"`
	BonusType generic.Nullable[int32] `json:"bonus_type" enum:"DropItemBonusType"`
}
