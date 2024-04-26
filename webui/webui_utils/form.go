package webui_utils

import (
	"elichika/utils"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SafeGetFormInt32(ctx *gin.Context, key string) (int32, bool) {
	form := ctx.MustGet("form").(*multipart.Form)
	_, exist := form.Value[key]
	if !exist {
		return 0, false
	}
	strValue := form.Value[key][0]
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return 0, false
	} else {
		return int32(intValue), true
	}
}

func GetFormInt32(ctx *gin.Context, key string) int32 {
	form := ctx.MustGet("form").(*multipart.Form)
	strValue := form.Value[key][0]
	intValue, err := strconv.Atoi(strValue)
	utils.CheckErr(err)
	return int32(intValue)
}

func GetFormString(ctx *gin.Context, key string) string {
	form := ctx.MustGet("form").(*multipart.Form)
	return form.Value[key][0]
}

func GetFormBool(ctx *gin.Context, key string) bool {
	form := ctx.MustGet("form").(*multipart.Form)
	onString, on := form.Value[key]
	if on && (onString[0] != "on") {
		panic("explicit off checkbox?")
	}
	return on
}
