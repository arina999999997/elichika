package request

import (
	"elichika/generic"
)

type CheerMemberGuildRequest struct {
	CheerItemAmount generic.Nullable[int32] `json:"cheer_item_amount"`
}
