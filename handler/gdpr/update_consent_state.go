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

	userId := int32(ctx.GetInt("user_id"))
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

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

	session.Finalize()
	common.JsonResponse(ctx, response.UpdateGdprConsentStateResponse{
		UserModel:     &session.UserModel,
		ConsentedInfo: loginData.GdprConsentedInfo,
	})
}

func init() {
	router.AddHandler("/gdpr/updateConsentState", updateConsentState)
}
