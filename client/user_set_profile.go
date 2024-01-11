package client

type UserSetProfile struct {
	UserSetProfileId        int32 `xorm:"pk 'user_set_profile_id'" json:"user_set_profile_id"`
	VoltageLiveDifficultyId int32 `xorm:"'voltage_live_difficulty_id'" json:"voltage_live_difficulty_id"`
	CommboLiveDifficultyId  int32 `xorm:"'commbo_live_difficulty_id'" json:"commbo_live_difficulty_id"`
}
