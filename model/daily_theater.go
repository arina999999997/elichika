package model

type UserDailyTheater struct {
	UserId         int  `xorm:"pk 'user_id'" json:"-"`
	DailyTheaterId int  `xorm:"pk 'daily_theater_id'" json:"daily_theater_id"`
	IsLiked        bool `xorm:"'is_liked'" json:"is_liked"`
}

func (udt *UserDailyTheater) Id() int64 {
	return int64(udt.DailyTheaterId)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_daily_theater"] = UserDailyTheater{}
}
