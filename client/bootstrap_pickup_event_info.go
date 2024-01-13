package client

import (
	"elichika/generic"
)

type BootstrapPickupEventInfo struct {
	EventId       int32                   `json:"event_id"`
	StartAt       int64                   `json:"start_at"`
	ClosedAt      int64                   `json:"closed_at"`
	EndAt         int64                   `json:"end_at"`
	EventType     int32                   `json:"event_type" enum:"EventType1"`
	BoosterItemId generic.Nullable[int32] `json:"booster_item_id"`
}
