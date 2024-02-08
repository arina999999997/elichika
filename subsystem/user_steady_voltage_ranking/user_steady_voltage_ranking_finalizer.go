package user_steady_voltage_ranking

import (
	"elichika/userdata"
	"elichika/utils"
)

func userSteadyVoltageRankingFinalizer(session *userdata.Session) {
	for _, userSteadyvoltageRanking := range session.UserModel.UserSteadyVoltageRankingById.Map {
		affected, err := session.Db.Table("u_steady_voltage_ranking").
			Where("user_id = ? AND steady_voltage_ranking_master_id = ?",
				session.UserId, userSteadyvoltageRanking.SteadyVoltageRankingMasterId).
			AllCols().Update(*userSteadyvoltageRanking)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_steady_voltage_ranking", *userSteadyvoltageRanking)
		}
	}
}

func init() {
	userdata.AddFinalizer(userSteadyVoltageRankingFinalizer)
}
