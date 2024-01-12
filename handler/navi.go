package handler

import (
	"elichika/config"
	"elichika/enum"
	"elichika/protocol/request"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// TODO(refactor): Change to use request and response types
func SaveUserNaviVoice(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveUserNaviVoiceRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	for _, naviVoiceMasterId := range req.NaviVoiceMasterIds {
		session.UpdateVoice(naviVoiceMasterId, false)
	}

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

// TODO(refactor): Change to use request and response types
func TapLovePoint(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	type TapLovePointReq struct {
		MemberMasterId int `json:"member_master_id"`
	}

	req := TapLovePointReq{}
	if err := json.Unmarshal([]byte(reqBody), &req); err != nil {
		panic(err)
	}
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	session.AddLovePoint(int32(req.MemberMasterId), int32(*config.Conf.TapBondGain))
	if session.UserStatus.TutorialPhase == enum.TutorialPhaseLovePointUp {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseTrainingLevelUp
	}
	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
