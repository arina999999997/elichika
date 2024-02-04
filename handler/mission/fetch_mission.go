package mission

import (
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"

	"github.com/gin-gonic/gin"
)

// TODO(mission): Implement stuff

func fetchMission(ctx *gin.Context) {
	// There is no request body
	userId := int32(ctx.GetInt("user_id"))
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

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/mission/fetchMission", fetchMission)
}
