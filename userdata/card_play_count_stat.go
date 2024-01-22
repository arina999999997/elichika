package userdata

import (
	"elichika/userdata/database"
	"elichika/utils"
)

func (session *Session) GetUserCardPlayCountStat(cardMasterId int32) database.UserCardPlayCountStat {
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

func (session *Session) UpdateUserCardPlayCountStat(userCardPlayCountStat database.UserCardPlayCountStat) {
	affected, err := session.Db.Table("u_card_play_count_stat").
		Where("user_id = ? AND card_master_id = ?", session.UserId, userCardPlayCountStat.CardMasterId).AllCols().Update(&userCardPlayCountStat)
	utils.CheckErr(err)
	if affected == 0 {
		genericDatabaseInsert(session, "u_card_play_count_stat", userCardPlayCountStat)
	}
}
