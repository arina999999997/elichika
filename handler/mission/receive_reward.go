package mission

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/log"
	"elichika/router"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func receiveReward(ctx *gin.Context) {
	req := request.MissionRewardRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	resp := user_mission.ReceiveReward(session, req.MissionIds.Slice)

	switch resp.(type) {
	case response.MissionReceiveResponse:
		common.JsonResponse(ctx, &resp)
	case response.MissionReceiveErrorResponse:
		common.AlternativeJsonResponse(ctx, &resp)
	default:
		log.Panic("not supported")
	}
}

func init() {
	router.AddHandler("/", "POST", "/mission/receiveReward", receiveReward)
}
