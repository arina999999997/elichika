package client

type LiveResultAchievementStatus struct {
	ClearCount       int32 `json:"clear_count"`
	GotVoltage       int32 `json:"got_voltage"`
	RemainingStamina int32 `json:"remaining_stamina"`
}
