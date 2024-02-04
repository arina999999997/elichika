package user_daily_theater

import (
	"elichika/userdata"
	"elichika/utils"
)

func dailyTheaterFinalizer(session *userdata.Session) {
	for _, userDailyTheater := range session.UserModel.UserDailyTheaterByDailyTheaterId.Map {
		affected, err := session.Db.Table("u_daily_theater").
			Where("user_id = ? AND daily_theater_id = ?",
				session.UserId, userDailyTheater.DailyTheaterId).
			AllCols().Update(userDailyTheater)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_daily_theater", *userDailyTheater)
		}
	}
}

func init() {
	userdata.AddFinalizer(dailyTheaterFinalizer)
}
