package request

import (
	"elichika/generic"
)

type RecoverLPRequest struct {
	ContentType int32                   `json:"content_type" enum:"ContentType"`
	ContentId   int32                   `json:"content_id"`
	Count       generic.Nullable[int32] `json:"count"`
}
