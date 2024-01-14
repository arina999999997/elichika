package request

import (
	"elichika/generic"
)

type GetPackUrlRequest struct {
	PackNames generic.List[string] `json:"pack_names"`
}
