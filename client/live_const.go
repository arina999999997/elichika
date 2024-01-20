package client

type LiveConst struct {
	// these fields actually have _<name>
	SpGaugeLength                int32 `json:"sp_gauge_length"`
	SpGaugeAdditionalRate        int32 `json:"sp_gauge_additional_rate"`
	SpGaugeReducingPoint         int32 `json:"sp_gauge_reducing_point"`
	SpSkillVoltageMagnification  int32 `json:"sp_skill_voltage_magnification"`
	NoteStaminaReduce            int32 `json:"note_stamina_reduce"`
	NoteVoltageUpperLimit        int32 `json:"note_voltage_upper_limit"`
	CollaboVoltageUpperLimit     int32 `json:"collabo_voltage_upper_limit"`
	SkillVoltageUpperLimit       int32 `json:"skill_voltage_upper_limit"`
	SquadChangeVoltageUpperLimit int32 `json:"squad_change_voltage_upper_limit"`
}
