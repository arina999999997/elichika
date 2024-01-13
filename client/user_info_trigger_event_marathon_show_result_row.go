package client

type UserInfoTriggerEventMarathonShowResultRow struct {
	TriggerId       int64 `json:"trigger_id"`
	EventMarathonId int32 `json:"event_marathon_id"`
	ResultAt        int64 `json:"result_at"`
	EndAt           int64 `json:"end_at"`
}
