package database

import (
	"elichika/client"
	"elichika/generic"

	"reflect"
)

func init() {
	AddTable("u_tower_card_used_count", generic.InterfaceWithAddedKey[int32](
		client.TowerCardUsedCount{},
		[]string{"UserId", "TowerId"},
		[]reflect.StructTag{`xorm:"pk 'user_id'"`, `xorm:"pk 'tower_id'"`},
	))
}
