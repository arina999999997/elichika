package user_tower

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func UpdateUserTowerVoltageRankingScore(session *userdata.Session, score database.UserTowerVoltageRankingScore) {
	affected, err := session.Db.Table("u_tower_voltage_ranking_score").
		Where("user_id = ? AND tower_id = ? AND floor_no = ?", session.UserId, score.TowerId, score.FloorNo).AllCols().
		Update(score)
	utils.CheckErr(err)
	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_tower_voltage_ranking_score", score)
	}
}
