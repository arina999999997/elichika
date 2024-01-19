package request

import (
	"elichika/generic"
)

type AddStoryLinkageRequest struct {
	CellId     int32                  `json:"cell_id"`
	IsAutoMode generic.Nullable[bool] `json:"is_auto_mode"`
}
