package client

type LiveNoteScore struct {
	JudgeType    int32 `json:"judge_type" enum:"JudgeType"`
	IsCritical   bool  `json:"is_critical"`
	Voltage      int32 `json:"voltage"`
	CardMasterId int32 `json:"card_master_id"`
	JudgedAt     int64 `json:"judged_at"`
}
