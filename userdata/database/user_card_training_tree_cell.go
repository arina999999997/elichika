package database

import (
	"elichika/client"
	"elichika/generic"

	"reflect"
)

type TrainingTreeCell struct {
	CardMasterId int   `xorm:"pk 'card_master_id'" json:"-"`
	CellId       int   `xorm:"pk 'cell_id'" json:"cell_id"`
	ActivatedAt  int64 `json:"activated_at"`
}

func init() {
	AddTable("u_card_training_tree_cell", generic.InterfaceWithAddedKey[int](
		client.UserCardTrainingTreeCell{},
		[]string{"UserId", "CardMasterId"},
		[]reflect.StructTag{`xorm:"'user_id'"`, `xorm:"'card_master_id'"`},
	))
}
