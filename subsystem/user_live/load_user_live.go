package user_live

import (
	"elichika/client"
	"elichika/userdata"
	"elichika/client/request"
	"elichika/utils"
)

func LoadUserLive(session *userdata.Session) (bool, client.Live, request.StartLiveRequest) {
	live := client.Live{}
	req := request.StartLiveRequest{}
	existLive, err := session.Db.Table("u_live").Where("user_id = ?", session.UserId).Get(&live)
	utils.CheckErr(err)
	existReq, err := session.Db.Table("u_start_live_request").Where("user_id = ?", session.UserId).Get(&req)
	utils.CheckErr(err)
	if !(existLive && existReq) {
		ClearUserLive(session)
		return false, live, req
	} else {
		return true, live, req
	}
}