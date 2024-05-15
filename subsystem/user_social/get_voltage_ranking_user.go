package user_social

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/utils"
)

func GetVoltageRankingUser(session *userdata.Session, userId int32) client.VoltageRankingUser {
	// TODO(cache): Maybe cache this
	user := client.VoltageRankingUser{}
	exist, err := session.Db.Table("u_status").Where("user_id = ?", userId).Cols(
		"user_id", "name", "rank", "recommend_card_master_id", "emblem_id").
		Get(&user.UserId, &user.UserName, &user.UserRank, &user.CardMasterId, &user.EmblemMasterId)
	utils.CheckErrMustExist(err, exist)
	exist, err = session.Db.Table("u_card").Where("user_id = ? AND card_master_id = ?", userId, user.CardMasterId).
		Cols("level", "is_awakening_image", "is_all_training_activated").
		Get(&user.Level, &user.IsAwakening, &user.IsAllTrainingActivated)
	utils.CheckErrMustExist(err, exist)
	return user
}
