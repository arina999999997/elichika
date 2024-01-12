package handler

import (
	"elichika/config"
	"elichika/userdata"

	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO(refactor): Change to use request and response types
func Agreement(ctx *gin.Context) {
	userId := ctx.GetInt("user_id")
	session := userdata.GetSession(ctx, userId)
	defer session.Close()

	signBody := session.Finalize("{}", "user_model")
	resp := SignResp(ctx, signBody, config.SessionKey)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
