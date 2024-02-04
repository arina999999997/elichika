package userdata

import (
	"elichika/client"
	"elichika/client/request"
	"elichika/utils"
)

// TODO(refactor): Move into subsystem
func (session *Session) SaveUserLive(live client.Live, req request.StartLiveRequest) {
	// delete whatever is there
	_, err := session.Db.Table("u_live").Where("user_id = ?", session.UserId).Delete(&client.Live{})
	utils.CheckErr(err)
	GenericDatabaseInsert(session, "u_live", live)
	_, err = session.Db.Table("u_start_live_request").Where("user_id = ?", session.UserId).Delete(&request.StartLiveRequest{})
	utils.CheckErr(err)
	GenericDatabaseInsert(session, "u_start_live_request", req)
}

func (session *Session) LoadUserLive() (bool, client.Live, request.StartLiveRequest) {
	live := client.Live{}
	req := request.StartLiveRequest{}
	existLive, err := session.Db.Table("u_live").Where("user_id = ?", session.UserId).Get(&live)
	utils.CheckErr(err)
	existReq, err := session.Db.Table("u_start_live_request").Where("user_id = ?", session.UserId).Get(&req)
	utils.CheckErr(err)
	if !(existLive && existReq) {
		session.ClearUserLive()
		return false, live, req
	} else {
		return true, live, req
	}
}

func (session *Session) ClearUserLive() {
	_, err := session.Db.Table("u_live").Where("user_id = ?", session.UserId).Delete(&client.Live{})
	utils.CheckErr(err)
	_, err = session.Db.Table("u_start_live_request").Where("user_id = ?", session.UserId).Delete(&request.StartLiveRequest{})
	utils.CheckErr(err)
}
