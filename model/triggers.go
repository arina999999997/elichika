package model

type TriggerReadReq struct {
	TriggerId int64 `json:"trigger_id"` // same for all trigger, for now
}
