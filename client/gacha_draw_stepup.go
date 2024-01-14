package client

import (
	"elichika/generic"
)

type GachaDrawStepup struct {
	CurrentStep   int32                   `json:"current_step"`
	LoopCount     int32                   `json:"loop_count"`
	MaxLoop       generic.Nullable[int32] `json:"max_loop"`
	MaxStep       int32                   `json:"max_step"`
	IsMaxNextStep bool                    `json:"is_max_next_step"`
}
