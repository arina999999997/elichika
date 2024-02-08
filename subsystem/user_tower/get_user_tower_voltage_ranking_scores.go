package user_tower

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func GetUserTowerVoltageRankingScores(session *userdata.Session, towerId int32) []database.UserTowerVoltageRankingScore {
	scores := []database.UserTowerVoltageRankingScore{}
	err := session.Db.Table("u_tower_voltage_ranking_score").
		Where("user_id = ? AND tower_id = ?", session.UserId, towerId).OrderBy("floor_no").Find(&scores)
	utils.CheckErr(err)
	return scores
}
