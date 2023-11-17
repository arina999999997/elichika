package handler

import (
	"elichika/config"
	"elichika/userdata"
	"elichika/utils"
	"elichika/protocol/request"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func SaveUserNaviVoice(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveUserNaviVoiceRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)
	
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()

	for _, naviVoiceMasterID := range req.NaviVoiceMasterIDs {
		session.UpdateVoice(naviVoiceMasterID, false)
	}

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func TapLovePoint(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type TapLovePointReq struct {
		MemberMasterID int `json:"member_master_id"`
	}

	req := TapLovePointReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}
	userID := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userID)
	defer session.Close()
	session.AddLovePoint(req.MemberMasterID, config.Conf.TapBondGain)
	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
