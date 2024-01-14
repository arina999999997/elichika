package handler

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/locale"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func GetPackUrl(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0].String()
	req := request.GetPackUrlRequest{}
	err := json.Unmarshal([]byte(reqBody), &req)
	utils.CheckErr(err)

	// these are hardcoded links to the original asset version
	cdnMasterVersion := "2d61e7b4e89961c7"
	if ctx.MustGet("locale").(*locale.Locale).MasterVersion == config.MasterVersionJp {
		cdnMasterVersion = "b66ec2295e9a00aa"
	}

	resp := response.GetPackUrlResponse{}
	for _, pack := range req.PackNames.Slice {
		resp.UrlList.Append(*config.Conf.CdnServer + "/" + cdnMasterVersion + "/" + pack)
	}

	JsonResponse(ctx, &resp)
}
