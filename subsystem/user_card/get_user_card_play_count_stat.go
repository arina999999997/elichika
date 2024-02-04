package user_card

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func GetUserCardPlayCountStat(session* userdata.Session, cardMasterId int32) database.UserCardPlayCountStat {
	res := database.UserCardPlayCountStat{}
	exist, err := session.Db.Table("u_card_play_count_stat").
		Where("user_id = ? AND card_master_id = ?", session.UserId, cardMasterId).Get(&res)
	utils.CheckErr(err)
	if !exist {
		res = database.UserCardPlayCountStat{
			CardMasterId: cardMasterId,
		}
	}
	return res
}
