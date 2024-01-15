package handler

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func UpdateConsentState(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.UpdateGdprConsentStateRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	session.UserStatus.GdprVersion = req.Version
	loginData := session.GetLoginData()
	for _, consent := range req.ConsentList.Slice {
		switch consent.GdprType {
		case enum.GdprConsentTypeAdIdIos:
			fallthrough
		case enum.GdprConsentTypeAdIdAndroid:
			fallthrough
		case enum.GdprConsentTypePersonalizedAd:
			loginData.GdprConsentedInfo.HasConsentedAdPurposeOfUse = consent.HasConsented
		case enum.GdprConsentTypeCrashReport:
			loginData.GdprConsentedInfo.HasConsentedCrashReport = consent.HasConsented
		}
	}
	session.UpdateLoginData(loginData)

	session.Finalize("{}", "dummy")
	JsonResponse(ctx, response.UpdateGdprConsentStateResponse{
		UserModel:     &session.UserModel,
		ConsentedInfo: loginData.GdprConsentedInfo,
	})
}
