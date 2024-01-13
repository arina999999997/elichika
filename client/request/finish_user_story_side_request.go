package request

import (
	"elichika/generic"
)

type FinishUserStorySideRequest struct {
	StorySideMasterId int32                  `json:"story_side_master_id"`
	IsAutoMode        generic.Nullable[bool] `json:"is_auto_mode"`
}
