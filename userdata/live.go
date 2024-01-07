package userdata

import (
	"elichika/model"
	"elichika/utils"
)

func (session *Session) SaveUserLive(live model.UserLive) {
	// delete whatever is there
	_, err := session.Db.Table("u_live_state").Where("user_id = ?", session.UserId).Delete(&model.UserLive{})
	utils.CheckErr(err)
	genericDatabaseInsert(session, "u_live_state", live)
}

func (session *Session) LoadUserLive() (bool, model.UserLive) {
	live := model.UserLive{}
	exist, err := session.Db.Table("u_live_state").Where("user_id = ?", session.UserId).Get(&live)
	utils.CheckErr(err)
	if exist {
		_, err = session.Db.Table("u_live_state").Where("user_id = ?", session.UserId).Delete(&model.UserLive{})
		utils.CheckErr(err)
	}
	return exist, live
}
