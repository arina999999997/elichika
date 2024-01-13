package handler

import (
	"elichika/config"
	"elichika/encrypt"
	"elichika/locale"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	presetDataPath = "presets/"
)

func SignResp(ctx *gin.Context, body, key string) (resp string) {
	ep := ctx.MustGet("ep").(string)
	masterVersion := ctx.MustGet("locale").(*locale.Locale).MasterVersion
	signBody := fmt.Sprintf("%d,\"%s\",0,%s", time.Now().UnixMilli(), masterVersion, body)
	sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+signBody), []byte(key))

	resp = fmt.Sprintf("[%s,\"%s\"]", signBody, sign)
	return
}

func JsonResponse(ctx *gin.Context, resp any) {
	signBody, err := json.Marshal(resp)
	// fmt.Println(string(signBody))
	utils.CheckErr(err)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, SignResp(ctx, string(signBody), config.SessionKey))
}

func GetData(fileName string) string {
	presetDataFile := presetDataPath + fileName
	if !utils.PathExists(presetDataFile) {
		panic("File not exist: " + fileName)
	}

	return utils.ReadAllText(presetDataFile)
}
