package user_login_bonus

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func updateUserLoginBonus(session *userdata.Session, userLoginBonus database.UserLoginBonus) {
	affected, err := session.Db.Table("u_login_bonus").
		Where("user_id = ? AND login_bonus_id = ?", session.UserId, userLoginBonus.LoginBonusId).
		AllCols().Update(userLoginBonus)
	utils.CheckErr(err)
	if affected == 0 {
		userdata.GenericDatabaseInsert(session, "u_login_bonus", userLoginBonus)
	}
}
