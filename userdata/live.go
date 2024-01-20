package userdata

import (
	"elichika/client"
	"elichika/utils"
)

func (session *Session) SaveUserLive(live client.Live) {
	// delete whatever is there
	_, err := session.Db.Table("u_live_state").Where("user_id = ?", session.UserId).Delete(&client.Live{})
	utils.CheckErr(err)
	genericDatabaseInsert(session, "u_live_state", live)
}

func (session *Session) LoadUserLive() (bool, client.Live) {
	live := client.Live{}
	exist, err := session.Db.Table("u_live_state").Where("user_id = ?", session.UserId).Get(&live)
	utils.CheckErr(err)
	if exist {
		_, err = session.Db.Table("u_live_state").Where("user_id = ?", session.UserId).Delete(&client.Live{})
		utils.CheckErr(err)
	}
	return exist, live
}
