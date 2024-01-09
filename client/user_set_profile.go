package client

type UserSetProfile struct {
	UserSetProfileId        int32 `xorm:"-" json:"user_set_profile_id"` // always 0
	VoltageLiveDifficultyId int32 `xorm:"'voltage_live_difficulty_id'" json:"voltage_live_difficulty_id"`
	CommboLiveDifficultyId  int32 `xorm:"'commbo_live_difficulty_id'" json:"commbo_live_difficulty_id"`
}

func (usp *UserSetProfile) Id() int64 {
	return int64(usp.UserSetProfileId)
}
