package handler

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/handler/common"
	"elichika/router"
	"elichika/subsystem/user_rule_description"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func saveRuleDescription(ctx *gin.Context) {
	req := request.SaveRuleDescriptionRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	// response with user model
	session := ctx.MustGet("session").(*userdata.Session)

	for _, ruleDescriptionId := range req.RuleDescriptionMasterIds.Slice {
		user_rule_description.UpdateUserRuleDescription(session, ruleDescriptionId)
	}

	common.JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func init() {
	router.AddHandler("/ruleDescription/saveRuleDescription", saveRuleDescription)
}
