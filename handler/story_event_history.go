package handler

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/item"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func UnlockStory(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UnlockStoryEventHistoryRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UnlockEventStory(req.EventStoryMasterId)
	session.RemoveResource(item.MemoryKey)

	JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func FinishStory(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.FinishStoryEventHistoryRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	// there is no cleared tracking so all this request does is set story mode
	if req.IsAutoMode.HasValue {
		session.UserStatus.IsAutoMode = req.IsAutoMode.Value
	}

	JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}
