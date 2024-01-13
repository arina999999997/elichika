package client

type UserInfoTriggerEventMiningShowResultRow struct {
	TriggerId     int64 `json:"trigger_id"`
	EventMiningId int32 `json:"event_mining_id"`
	ResultAt      int64 `json:"result_at"`
	EndAt         int64 `json:"end_at"`
}
