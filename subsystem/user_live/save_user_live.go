package user_live

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/client/request"
	"elichika/utils"
)

func SaveUserLive(session *userdata.Session, live client.Live, req request.StartLiveRequest) {
	// delete whatever is there
	_, err := session.Db.Table("u_live").Where("user_id = ?", session.UserId).Delete(&client.Live{})
	utils.CheckErr(err)
	userdata.GenericDatabaseInsert(session, "u_live", live)
	_, err = session.Db.Table("u_start_live_request").Where("user_id = ?", session.UserId).Delete(&request.StartLiveRequest{})
	utils.CheckErr(err)
	userdata.GenericDatabaseInsert(session, "u_start_live_request", req)
}