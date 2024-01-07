package userdata

import (
	"elichika/utils"
)

func dailyTheaterFinalizer(session *Session) {
	for _, userDailyTheater := range session.UserModel.UserDailyTheaterByDailyTheaterId.Objects {
		affected, err := session.Db.Table("u_daily_theater").
			Where("user_id = ? AND daily_theater_id = ?",
				session.UserStatus.UserId, userDailyTheater.DailyTheaterId).
			AllCols().Update(userDailyTheater)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_daily_theater").
				Insert(userDailyTheater)
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(dailyTheaterFinalizer)
	addGenericTableFieldPopulator("u_daily_theater", "UserDailyTheaterByDailyTheaterId")
}
