package asset

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/handler/common"
	"elichika/locale"
	"elichika/router"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func getPackUrl(ctx *gin.Context) {
	req := request.GetPackUrlRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	// these are hardcoded links to the original asset version
	cdnMasterVersion := "2d61e7b4e89961c7"
	if ctx.MustGet("locale").(*locale.Locale).MasterVersion == config.MasterVersionJp {
		cdnMasterVersion = "b66ec2295e9a00aa"
	}

	resp := response.GetPackUrlResponse{}
	for _, pack := range req.PackNames.Slice {
		if *config.Conf.CdnServer == "http://127.0.0.1:8080/static" {
			resp.UrlList.Append(*config.Conf.CdnServer + "/" + "assets" + "/" + pack)
		} else {
			resp.UrlList.Append(*config.Conf.CdnServer + "/" + cdnMasterVersion + "/" + pack)
		}
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/asset/getPackUrl", getPackUrl)
}
