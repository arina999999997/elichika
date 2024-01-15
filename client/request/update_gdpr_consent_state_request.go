package request

import (
	"elichika/client"
	"elichika/generic"
)

type UpdateGdprConsentStateRequest struct {
	Version     int32                                `json:"version"`
	ConsentList generic.List[client.GdprConsentInfo] `json:"consent_list"`
}
