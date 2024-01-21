package info_trigger

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
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

	common.JsonResponse(ctx, &response.UserModelResponse{
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

	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func TriggerReadMemberLoveLevelUp(ctx *gin.Context) {
	// there is no request body
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.ReadAllMemberLoveLevelUpTriggers()

	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func TriggerReadMemberGuildSupportItemExpired(ctx *gin.Context) {
	// there is no request body
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.ReadMemberGuildSupportItemExpired()

	common.JsonResponse(ctx, &response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/infoTrigger/read", TriggerRead)
	router.AddHandler("/infoTrigger/readCardGradeUp", TriggerReadCardGradeUp)
	router.AddHandler("/infoTrigger/readMemberLoveLevelUp", TriggerReadMemberLoveLevelUp)
	router.AddHandler("/infoTrigger/readMemberGuildSupportItemExpired", TriggerReadMemberGuildSupportItemExpired)
}
