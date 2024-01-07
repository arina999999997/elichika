package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) GetUserLoginBonus(loginBonusId int) model.UserLoginBonus {
	userLoginBonus := model.UserLoginBonus{}
	exists, err := session.Db.Table("u_login_bonus").
		Where("user_id = ? AND login_bonus_id = ?", session.UserStatus.UserId, loginBonusId).Get(&userLoginBonus)
	utils.CheckErr(err)
	if !exists {
		userLoginBonus = model.UserLoginBonus{
			UserId:             session.UserStatus.UserId,
			LoginBonusId:       loginBonusId,
			LastReceivedReward: -1,
			LastReceivedAt:     0,
		}
	}
	return userLoginBonus
}

func (session *Session) UpdateUserLoginBonus(userLoginBonus model.UserLoginBonus) {
	affected, err := session.Db.Table("u_login_bonus").
		Where("user_id = ? AND login_bonus_id = ?", session.UserStatus.UserId, userLoginBonus.LoginBonusId).
		AllCols().Update(userLoginBonus)
	utils.CheckErr(err)
	if affected == 0 {
		_, err := session.Db.Table("u_login_bonus").Insert(userLoginBonus)
		utils.CheckErr(err)
	}
}
