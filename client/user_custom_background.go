package client

type UserCustomBackground struct {
	CustomBackgroundMasterId int32 `xorm:"pk 'custom_background_master_id'" json:"custom_background_master_id"`
	IsNew                    bool  `xorm:"'is_new'" json:"is_new"`
}

func (ucb *UserCustomBackground) Id() int64 {
	return int64(ucb.CustomBackgroundMasterId)
}
