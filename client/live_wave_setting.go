package client

type LiveWaveSetting struct {
	Id            int32 `json:"id"`
	WaveDamage    int32 `json:"wave_damage"`
	MissionType   int32 `json:"mission_type" enum:"LiveAppealTimeMission"`
	Arg1          int32 `json:"arg_1"`
	Arg2          int32 `json:"arg_2"`
	RewardVoltage int32 `json:"reward_voltage"`
}
