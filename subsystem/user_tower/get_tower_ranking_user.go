package user_tower

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetTowerRankingUser(session *userdata.Session, rankingUserId int32) client.TowerRankingUser {
	user := client.TowerRankingUser{}

	exist, err := session.Db.Table("u_status").Where("user_id = ?", rankingUserId).Cols(
		"user_id", "name", "rank", "recommend_card_master_id", "emblem_id").Get(
		&user.UserId, &user.UserName, &user.UserRank, &user.CardMasterId, &user.EmblemMasterId)
	utils.CheckErrMustExist(err, exist)

	exist, err = session.Db.Table("u_card").Where("user_id = ? AND card_master_id = ?", rankingUserId, user.CardMasterId).Cols(
		"level", "is_awakening_image", "is_all_training_activated").Get(
		&user.Level, &user.IsAwakening, &user.IsAllTrainingActivated)
	utils.CheckErrMustExist(err, exist)

	return user
}
