package model

import (
	"elichika/generic"
)

type TrainingTreeCell struct {
	CardMasterId int   `xorm:"pk 'card_master_id'" json:"-"`
	CellId       int   `xorm:"pk 'cell_id'" json:"cell_id"`
	ActivatedAt  int64 `json:"activated_at"` // int64 so we don't have Y2K38 problem
}

func init() {

	TableNameToInterface["u_training_tree_cell"] = generic.UserIdWrapper[TrainingTreeCell]{}
}
