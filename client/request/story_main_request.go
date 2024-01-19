package request

import (
	"elichika/generic"
)

type StoryMainRequest struct {
	CellId     int32                   `json:"cell_id"`
	IsAutoMode generic.Nullable[bool]  `json:"is_auto_mode"`
	MemberId   generic.Nullable[int32] `json:"member_id"`
}
