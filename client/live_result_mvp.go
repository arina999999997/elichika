package client

type LiveResultMvp struct {
	CardMasterId        int32 `json:"card_master_id"`
	GetVoltage          int32 `json:"get_voltage"`
	SkillTriggeredCount int32 `json:"skill_triggered_count"`
	AppealCount         int32 `json:"appeal_count"`
}
