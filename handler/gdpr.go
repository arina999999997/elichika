package handler

import (
	"elichika/userdata"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// TODO(refactor): Change to use request and response types
func UpdateConsentState(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()
	session.UserStatus.GdprVersion = int32(reqBody.Get("version").Int())
	session.Finalize("{}", "user_model")
	// Don't know the format of this response, but we can set gdpr version to 4 to skip it.
	// TODO(very_low): read the decompiled code and see what the format is
	// resp := SignResp(ctx, signBody, config.SessionKey)
	// ctx.Header("Content-Type", "application/json")
	// ctx.String(500, resp)
}
