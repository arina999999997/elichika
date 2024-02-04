package user_card

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func UpdateUserCardPlayCountStat(session *userdata.Session, userCardPlayCountStat database.UserCardPlayCountStat) {
	affected, err := session.Db.Table("u_card_play_count_stat").
		Where("user_id = ? AND card_master_id = ?", session.UserId, userCardPlayCountStat.CardMasterId).AllCols().Update(&userCardPlayCountStat)
	utils.CheckErr(err)
	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_card_play_count_stat", userCardPlayCountStat)
	}
}
