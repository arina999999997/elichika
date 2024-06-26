package common

import (
	"elichika/encrypt"
	"elichika/locale"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SignResp(ctx *gin.Context, body string, key []byte) (resp string) {
	ep := ctx.MustGet("ep").(string)
	masterVersion := ctx.MustGet("locale").(*locale.Locale).MasterVersion
	signBody := fmt.Sprintf("%d,\"%s\",0,%s", time.Now().UnixMilli(), masterVersion, body)
	sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+signBody), []byte(key))

	resp = fmt.Sprintf("[%s,\"%s\"]", signBody, sign)
	return
}

func JsonResponse(ctx *gin.Context, resp any) {
	session := ctx.MustGet("session").(*userdata.Session)
	if session != nil {
		session.Finalize() // calling this multiple time is fine, sometime we want some result that is only obtained after finalizing
	}
	signBody, err := json.Marshal(resp)
	// fmt.Println(string(signBody))
	utils.CheckErr(err)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, SignResp(ctx, string(signBody), ctx.MustGet("sign_key").([]byte)))
}

func SignRespWithRespnoseType(ctx *gin.Context, body string, key []byte, rType int32) (resp string) {
	ep := ctx.MustGet("ep").(string)
	masterVersion := ctx.MustGet("locale").(*locale.Locale).MasterVersion
	signBody := fmt.Sprintf("%d,\"%s\",%d,%s", time.Now().UnixMilli(), masterVersion, rType, body)
	sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+signBody), []byte(key))

	resp = fmt.Sprintf("[%s,\"%s\"]", signBody, sign)
	return
}

func AlternativeJsonResponse(ctx *gin.Context, resp any) {
	session := ctx.MustGet("session").(*userdata.Session)
	if session != nil {
		session.Finalize() // calling this multiple time is fine
	}
	signBody, err := json.Marshal(resp)
	// fmt.Println(string(signBody))
	// fmt.Println(ctx.MustGet("sign_key").([]byte))
	utils.CheckErr(err)
	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, SignRespWithRespnoseType(ctx, string(signBody), ctx.MustGet("sign_key").([]byte), 1))
}
