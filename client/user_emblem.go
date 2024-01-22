package client

import (
	"elichika/generic"
)

type UserEmblem struct {
	EmblemMId   int32                    `xorm:"pk 'emblem_m_id'" json:"emblem_m_id"`
	IsNew       bool                     `xorm:"'is_new'" json:"is_new"`
	EmblemParam generic.Nullable[string] `xorm:"json 'emblem_param'" json:"emblem_param"`
	AcquiredAt  int64                    `xorm:"'acquired_at'" json:"acquired_at"`
}
