package handler

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func SaveUserNaviVoice(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.SaveUserNaviVoiceRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	for _, naviVoiceMasterId := range req.NaviVoiceMasterIds.Slice {
		session.UpdateVoice(naviVoiceMasterId, false)
	}

	session.Finalize()
	JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}

func TapLovePoint(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.TapLovePointRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.AddLovePoint(req.MemberMasterId, *config.Conf.TapBondGain)
	if session.UserStatus.TutorialPhase == enum.TutorialPhaseLovePointUp {
		session.UserStatus.TutorialPhase = enum.TutorialPhaseTrainingLevelUp
	}

	session.Finalize()
	JsonResponse(ctx, response.UserModelResponse{
		UserModel: &session.UserModel,
	})
}
