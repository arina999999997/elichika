package handler

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// TODO(mission): Implement stuff

func FetchMission(ctx *gin.Context) {
	// There is no request body
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	resp := response.FetchMissionResponse{
		UserModel: &session.UserModel,
	}
	// this is the official server behaviour
	session.PopulateUserModelField("UserMissionByMissionId")
	for _, mission := range session.UserModel.UserMissionByMissionId.Map {
		resp.MissionMasterIdList.Append(mission.MissionMId)
	}

	JsonResponse(ctx, &resp)
}

func ClearMissionBadge(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ClearMissionNewBadgeRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}
