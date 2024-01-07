package handler

import (
	"elichika/config"
	"elichika/userdata"
	"elichika/utils"
	// "elichika/enum"
	"elichika/protocol/request"

	"encoding/json"
	"net/http"

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

	for _, ruleDescriptionId := range req.RuleDescriptionMasterIds {
		session.UpdateUserRuleDescription(ruleDescriptionId)
	}

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
