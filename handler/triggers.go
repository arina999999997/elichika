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

func TriggerReadCardGradeUp(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ReadInfoTriggerRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	session.RemoveTriggerCardGradeUp(req.TriggerId)

	JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func TriggerRead(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.ReadInfoTriggerRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	session.DeleteTriggerBasic(req.TriggerId)

	JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func TriggerReadMemberLoveLevelUp(ctx *gin.Context) {
	// there is no request body
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.ReadAllMemberLoveLevelUpTriggers()

	JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func TriggerReadMemberGuildSupportItemExpired(ctx *gin.Context) {
	// there is no request body
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.ReadMemberGuildSupportItemExpired()

	JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}
