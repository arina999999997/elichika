package voltage_ranking

import (
	"elichika/userdata"
	"elichika/utils"
)

// TODO(storage): remove non-top record or we will have lots of storage
func SetVoltageRanking(session *userdata.Session, userVoltageRanking UserVoltageRanking) {
	GetRankingByLiveDifficultyId(session, userVoltageRanking.LiveDifficultyId).
		Update(userVoltageRanking.UserId, userVoltageRanking.VoltagePoint)
	affected, err := session.Db.Table("u_voltage_ranking").
		Where("user_id = ? AND live_difficulty_id = ?", userVoltageRanking.UserId, userVoltageRanking.LiveDifficultyId).
		AllCols().Update(&userVoltageRanking)
	utils.CheckErr(err)
	if affected == 0 {
		_, err = session.Db.Table("u_voltage_ranking").Insert(&userVoltageRanking)
		utils.CheckErr(err)
	}
}
