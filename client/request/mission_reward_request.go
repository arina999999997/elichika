package request

import (
	"elichika/generic"
)

type MissionRewardRequest struct {
	MissionIds generic.Array[int32] `json:"mission_ids"`
}
