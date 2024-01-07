package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) SaveUserLive(live model.UserLive) {
	// delete whatever is there
	_, err := session.Db.Table("u_live_state").Where("user_id = ?", live.UserId).Delete(&model.UserLive{})
	utils.CheckErr(err)
	affected, err := session.Db.Table("u_live_state").AllCols().Insert(live)
	utils.CheckErr(err)
	if affected != 1 {
		panic("failed to insert")
	}
}

func (session *Session) LoadUserLive() (bool, model.UserLive) {
	live := model.UserLive{}
	exist, err := session.Db.Table("u_live_state").Where("user_id = ?", session.UserStatus.UserId).Get(&live)
	utils.CheckErr(err)
	if exist {
		_, err = session.Db.Table("u_live_state").Where("user_id = ?", session.UserStatus.UserId).Delete(&model.UserLive{})
		utils.CheckErr(err)
	}
	return exist, live
}
