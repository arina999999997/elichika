package user_suit

import (
	"elichika/userdata"
	"elichika/utils"
)

func userSuitFinalizer(session *userdata.Session) {
	for _, suit := range session.UserModel.UserSuitBySuitId.Map {
		affected, err := session.Db.Table("u_suit").
			Where("user_id = ? AND suit_master_id = ?", session.UserId, suit.SuitMasterId).
			Update(*suit)
		utils.CheckErr(err)
		if affected == 0 {
			userdata.GenericDatabaseInsert(session, "u_suit", *suit)
		}
	}
}

func init() {
	userdata.AddContentFinalizer(userSuitFinalizer)
}
