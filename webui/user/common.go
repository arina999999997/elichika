package user

import (
	"elichika/utils"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Response *string `json:"response"`
	Error    *string `json:"error"`
}

func commonResponse(ctx *gin.Context, responseStr, errorStr string) {
	resp := Response{}
	if errorStr == "" {
		resp.Response = &responseStr
	} else {
		resp.Error = &errorStr
	}

	jsonBytes, err := json.Marshal(resp)
	utils.CheckErr(err)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(jsonBytes))
}
