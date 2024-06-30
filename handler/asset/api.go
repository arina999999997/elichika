package asset

import (
	"elichika/router"
	"elichika/utils"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func staticApi(ctx *gin.Context) {
	masterVersion, exist := ctx.GetQuery("master")
	utils.MustExist(exist)
	file, exist := ctx.GetQuery("file")
	utils.MustExist(exist)
	startString, exist := ctx.GetQuery("start")
	utils.MustExist(exist)
	start, err := strconv.Atoi(startString)
	utils.CheckErr(err)
	sizeString, exist := ctx.GetQuery("size")
	utils.MustExist(exist)
	size, err := strconv.Atoi(sizeString)
	utils.CheckErr(err)

	sendRange(ctx, fmt.Sprintf("static/%s/%s", masterVersion, file), start, size)
}

func init() {
	router.AddHandler("/static_api", "GET", "/", staticApi)
}
