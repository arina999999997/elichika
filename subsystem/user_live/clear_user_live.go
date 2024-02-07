package user_live

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/userdata"
	"elichika/utils"
)

func ClearUserLive(session *userdata.Session) {
	_, err := session.Db.Table("u_live").Where("user_id = ?", session.UserId).Delete(&client.Live{})
	utils.CheckErr(err)
	_, err = session.Db.Table("u_start_live_request").Where("user_id = ?", session.UserId).Delete(&request.StartLiveRequest{})
	utils.CheckErr(err)
}
