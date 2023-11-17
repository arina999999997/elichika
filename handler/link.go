package handler

import (
	"elichika/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Sif2DataLink(ctx *gin.Context) {
	signBody := GetData("sif2DataLink.json")
	resp := SignResp(ctx, signBody, config.SessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}
