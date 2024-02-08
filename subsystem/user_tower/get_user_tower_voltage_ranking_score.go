package user_tower

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func GetUserTowerVoltageRankingScore(session *userdata.Session, towerId, floorNo int32) database.UserTowerVoltageRankingScore {
	score := database.UserTowerVoltageRankingScore{}
	exists, err := session.Db.Table("u_tower_voltage_ranking_score").
		Where("user_id = ? AND tower_id = ? AND floor_no = ?", session.UserId, towerId, floorNo).Get(&score)
	utils.CheckErr(err)
	if !exists {
		score = database.UserTowerVoltageRankingScore{
			TowerId: towerId,
			FloorNo: floorNo,
			Voltage: 0,
		}
	}
	return score
}
