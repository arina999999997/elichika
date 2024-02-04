package userdata

import (
	"elichika/userdata/database"
	"elichika/utils"
)

// TODO(refactor): Move into subsystem
func (session *Session) GetUserLoginBonus(loginBonusId int32) database.UserLoginBonus {
	userLoginBonus := database.UserLoginBonus{}
	exists, err := session.Db.Table("u_login_bonus").
		Where("user_id = ? AND login_bonus_id = ?", session.UserId, loginBonusId).Get(&userLoginBonus)
	utils.CheckErr(err)
	if !exists {
		userLoginBonus = database.UserLoginBonus{
			LoginBonusId:       loginBonusId,
			LastReceivedReward: -1,
			LastReceivedAt:     0,
		}
	}
	return userLoginBonus
}

func (session *Session) UpdateUserLoginBonus(userLoginBonus database.UserLoginBonus) {
	affected, err := session.Db.Table("u_login_bonus").
		Where("user_id = ? AND login_bonus_id = ?", session.UserId, userLoginBonus.LoginBonusId).
		AllCols().Update(userLoginBonus)
	utils.CheckErr(err)
	if affected == 0 {
		GenericDatabaseInsert(session, "u_login_bonus", userLoginBonus)
	}
}
