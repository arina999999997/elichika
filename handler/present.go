package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO(refactor): Change to use request and response types
func FetchPresent(ctx *gin.Context) {
	signBody := GetData("fetchPresent.json")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
