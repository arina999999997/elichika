package userdata

import (
	"elichika/utils"
)

func missionFinalizer(session *Session) {
	for _, userMission := range session.UserModel.UserMissionByMissionId.Objects {
		affected, err := session.Db.Table("u_mission").Where("user_id = ? AND mission_m_id = ?",
			session.UserStatus.UserId, userMission.MissionMId).AllCols().Update(userMission)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_mission").Insert(userMission)
			utils.CheckErr(err)
		}
	}
	for _, userDailyMission := range session.UserModel.UserDailyMissionByMissionId.Objects {
		affected, err := session.Db.Table("u_daily_mission").Where("user_id = ? AND mission_m_id = ?",
			session.UserStatus.UserId, userDailyMission.MissionMId).AllCols().Update(userDailyMission)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_daily_mission").Insert(userDailyMission)
			utils.CheckErr(err)
		}
	}
	for _, userWeeklyMission := range session.UserModel.UserWeeklyMissionByMissionId.Objects {
		affected, err := session.Db.Table("u_weekly_mission").Where("user_id = ? AND mission_m_id = ?",
			session.UserStatus.UserId, userWeeklyMission.MissionMId).AllCols().Update(userWeeklyMission)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_weekly_mission").Insert(userWeeklyMission)
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(missionFinalizer)
	addGenericTableFieldPopulator("u_mission", "UserMissionByMissionId")
	addGenericTableFieldPopulator("u_daily_mission", "UserDailyMissionByMissionId")
	addGenericTableFieldPopulator("u_weekly_mission", "UserWeeklyMissionByMissionId")
}
