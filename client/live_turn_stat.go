package client

type LiveTurnStat struct {
	NoteId            int32 `json:"note_id"`
	CurrentLife       int32 `json:"current_life"`
	CurrentVoltage    int32 `json:"current_voltage"`
	AppendedShield    int32 `json:"appended_shield"`
	HealedLife        int32 `json:"healed_life"`
	HealedLifePercent int32 `json:"healed_life_percent"`
	StaminaDamage     int32 `json:"stamina_damage"`
}
