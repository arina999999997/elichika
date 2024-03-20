package mission

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_mission"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func receiveReward(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.MissionRewardRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := user_mission.ReceiveReward(session, req.MissionIds.Slice)
	session.Finalize()

	switch resp.(type) {
	case response.MissionReceiveResponse:
		common.JsonResponse(ctx, &resp)
	case response.MissionReceiveErrorResponse:
		common.JsonResponseWithRespnoseType(ctx, &resp, 1)
	default:
		panic("not supported")
	}
}

func init() {
	router.AddHandler("/mission/receiveReward", receiveReward)
}
