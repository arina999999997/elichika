package client

type LiveEventMarathonStatus struct {
	EventId                   int32 `json:"event_id"`
	IsUseEventMarathonBooster bool  `json:"is_use_event_marathon_booster"`
}
