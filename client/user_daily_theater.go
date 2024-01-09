package client

type UserDailyTheater struct {
	DailyTheaterId int32 `xorm:"pk 'daily_theater_id'" json:"daily_theater_id"`
	IsLiked        bool  `xorm:"'is_liked'" json:"is_liked"`
}

func (udt *UserDailyTheater) Id() int64 {
	return int64(udt.DailyTheaterId)
}
