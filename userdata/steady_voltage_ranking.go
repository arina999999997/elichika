package userdata

import (
	"elichika/utils"
)

func steadyVoltageRankingFinalizer(session *Session) {
	for _, userSteadyvoltageRanking := range session.UserModel.UserSteadyVoltageRankingById.Objects {
		affected, err := session.Db.Table("u_steady_voltage_ranking").
			Where("user_id = ? AND steady_voltage_ranking_master_id = ?",
				session.UserStatus.UserId, userSteadyvoltageRanking.SteadyVoltageRankingMasterId).
			AllCols().Update(userSteadyvoltageRanking)
		utils.CheckErr(err)
		if affected == 0 {
			_, err = session.Db.Table("u_steady_voltage_ranking").
				Insert(userSteadyvoltageRanking)
			utils.CheckErr(err)
		}
	}
}

func init() {
	addFinalizer(steadyVoltageRankingFinalizer)
	addGenericTableFieldPopulator("u_steady_voltage_ranking", "UserSteadyVoltageRankingById")
}
