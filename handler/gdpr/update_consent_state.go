package gdpr

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/enum"
	"elichika/handler/common"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func updateConsentState(ctx *gin.Context) {
	req := request.UpdateGdprConsentStateRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	session.UserStatus.GdprVersion = req.Version
	loginData := session.GetLoginResponse()
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

	common.JsonResponse(ctx, response.UpdateGdprConsentStateResponse{
		UserModel:     &session.UserModel,
		ConsentedInfo: loginData.GdprConsentedInfo,
	})
}

func init() {
	router.AddHandler("/", "POST", "/gdpr/updateConsentState", updateConsentState)
}
