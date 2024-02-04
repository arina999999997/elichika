package userdata

import (
	"elichika/utils"
)

// TODO(refactor): Move into subsystem
func missionFinalizer(session *Session) {
	for _, userMission := range session.UserModel.UserMissionByMissionId.Map {
		affected, err := session.Db.Table("u_mission").Where("user_id = ? AND mission_m_id = ?",
			session.UserId, userMission.MissionMId).AllCols().Update(*userMission)
		utils.CheckErr(err)
		if affected == 0 {
			GenericDatabaseInsert(session, "u_mission", *userMission)
		}
	}
	for _, userDailyMission := range session.UserModel.UserDailyMissionByMissionId.Map {
		affected, err := session.Db.Table("u_daily_mission").Where("user_id = ? AND mission_m_id = ?",
			session.UserId, userDailyMission.MissionMId).AllCols().Update(*userDailyMission)
		utils.CheckErr(err)
		if affected == 0 {
			GenericDatabaseInsert(session, "u_daily_mission", *userDailyMission)
		}
	}
	for _, userWeeklyMission := range session.UserModel.UserWeeklyMissionByMissionId.Map {
		affected, err := session.Db.Table("u_weekly_mission").Where("user_id = ? AND mission_m_id = ?",
			session.UserId, userWeeklyMission.MissionMId).AllCols().Update(*userWeeklyMission)
		utils.CheckErr(err)
		if affected == 0 {
			GenericDatabaseInsert(session, "u_weekly_mission", *userWeeklyMission)
		}
	}
}

func init() {
	AddFinalizer(missionFinalizer)
}
