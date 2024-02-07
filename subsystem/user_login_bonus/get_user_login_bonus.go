package user_login_bonus

import (
	"elichika/userdata"
	"elichika/userdata/database"
	"elichika/utils"
)

func getUserLoginBonus(session *userdata.Session, loginBonusId int32) database.UserLoginBonus {
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
