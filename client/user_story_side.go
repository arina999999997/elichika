package client

type UserStorySide struct {
	StorySideMasterId int32 `xorm:"pk 'story_side_master_id'" json:"story_side_master_id"`
	IsNew             bool  `xorm:"'is_new'" json:"is_new"`
	AcquiredAt        int64 `xorm:"'acquired_at'" json:"acquired_at"`
}

func (uss *UserStorySide) Id() int64 {
	return int64(uss.StorySideMasterId)
}
