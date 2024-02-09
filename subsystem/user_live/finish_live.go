package user_live

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"
)

func FinishLive(session *userdata.Session, req request.FinishLiveRequest) response.FinishLiveResponse {
	// this is pretty different for different type of live
	// for simplicity we just read the request and call different handlers, even though we might be able to save some extra work

	exist, live, startReq := LoadUserLive(session)
	utils.MustExist(exist)
	ClearUserLive(session)
	// TODO(lp): Remove LP here if we want that

	switch live.LiveType {
	case enum.LiveTypeManual:
		return liveTypeManualHandler(session, req, live, startReq)
	case enum.LiveTypeTower:
		return liveTypeTowerHandler(session, req, live, startReq)
	default:
		panic("not handled")
	}
}
