package handler

import (
	"elichika/config"
	"elichika/model"
	"elichika/userdata"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func TriggerReadCardGradeUp(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	UserID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, UserID)
	defer session.Close()
	req := model.TriggerReadReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	session.AddTriggerCardGradeUp(model.TriggerCardGradeUp{
		TriggerID: req.TriggerID,
		IsNull:    true,
	})
	resp := session.Finalize("{}", "user_model")
	resp = SignResp(ctx, resp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func TriggerRead(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	UserID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, UserID)
	defer session.Close()
	req := model.TriggerReadReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}

	session.AddTriggerBasic(model.TriggerBasic{
		TriggerID: req.TriggerID,
		IsNull:    true,
	})
	resp := session.Finalize("{}", "user_model")
	resp = SignResp(ctx, resp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func TriggerReadMemberLoveLevelUp(ctx *gin.Context) {
	// req is null, so we need to pull the triggers from db here
	UserID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, UserID)
	defer session.Close()

	triggers := session.GetAllTriggerMemberLoveLevelUps()
	for _, trigger := range triggers {
		trigger.IsNull = true
		session.AddTriggerMemberLoveLevelUp(trigger)
	}

	resp := session.Finalize("{}", "user_model")
	resp = SignResp(ctx, resp, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
