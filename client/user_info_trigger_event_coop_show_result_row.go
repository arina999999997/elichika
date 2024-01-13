package client

type UserInfoTriggerEventCoopShowResultRow struct {
	TriggerId     int64 `json:"trigger_id"`
	EventMasterId int32 `json:"event_master_id"`
	ResultAt      int64 `json:"result_at"`
	EndAt         int64 `json:"end_at"`
}
