package handler

import (
	"elichika/config"
	"elichika/model"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func TriggerReadCardGradeUp(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := model.TriggerReadReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.AddTriggerCardGradeUp(model.TriggerCardGradeUp{
		TriggerId: req.TriggerId,
		IsNull:    true,
	})

	resp := session.Finalize("{}", "user_model")
	resp = SignResp(ctx, resp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func TriggerRead(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := model.TriggerReadReq{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.AddTriggerBasic(model.TriggerBasic{
		TriggerId: req.TriggerId,
		IsNull:    true,
	})

	resp := session.Finalize("{}", "user_model")
	resp = SignResp(ctx, resp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func TriggerReadMemberLoveLevelUp(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.ReadAllMemberLoveLevelUpTriggers()

	resp := session.Finalize("{}", "user_model")
	resp = SignResp(ctx, resp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func TriggerReadMemberGuildSupportItemExpired(ctx *gin.Context) {
	// there's no body, fetch the trigger from db and remove it

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.ReadMemberGuildSupportItemExpired()

	resp := session.Finalize("{}", "user_model")
	resp = SignResp(ctx, resp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
