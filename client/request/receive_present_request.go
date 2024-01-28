package request

import (
	"elichika/generic"
)

type ReceivePresentRequest struct {
	Ids generic.List[int32] `json:"ids"`
}
