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

func SaveRuleDescription(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveRuleDescriptionRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	// response with user model
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	for _, ruleDescriptionId := range req.RuleDescriptionMasterIds.Slice {
		session.UpdateUserRuleDescription(ruleDescriptionId)
	}

	session.Finalize("{}", "dummy")
	JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}
