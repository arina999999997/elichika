package handler

import (
	"elichika/encrypt"
	"elichika/locale"
	"elichika/utils"

	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	presetDataPath = "assets/preset/"
)

func SignResp(ctx *gin.Context, body, key string) (resp string) {
	ep := ctx.MustGet("ep").(string)
	masterVersion := ctx.MustGet("locale").(*locale.Locale).MasterVersion
	signBody := fmt.Sprintf("%d,\"%s\",0,%s", time.Now().UnixMilli(), masterVersion, body)
	sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+signBody), []byte(key))

	resp = fmt.Sprintf("[%s,\"%s\"]", signBody, sign)
	return
}

func GetData(fileName string) string {
	presetDataFile := presetDataPath + fileName
	if !utils.PathExists(presetDataFile) {
		panic("File not exist: " + fileName)
	}

	return utils.ReadAllText(presetDataFile)
}
