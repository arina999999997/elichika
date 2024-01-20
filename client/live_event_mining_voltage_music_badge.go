package client

import (
	"elichika/generic"
)

type LiveEventMiningVoltageMusicBadge struct {
	LiveMasterId generic.Nullable[int32] `json:"live_master_id"`
	AppealText   LocalizedText           `json:"appeal_text"`
	EndAt        int64                   `json:"end_at"`
}
