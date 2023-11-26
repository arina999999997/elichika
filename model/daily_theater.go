package model

type UserDailyTheater struct {
	UserID         int  `xorm:"pk 'user_id'" json:"-"`
	DailyTheaterID int  `xorm:"pk 'daily_theater_id'" json:"daily_theater_id"`
	IsLiked        bool `xorm:"'is_liked'" json:"is_liked"`
}

func (udt *UserDailyTheater) ID() int64 {
	return int64(udt.DailyTheaterID)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_daily_theater"] = UserDailyTheater{}
}
