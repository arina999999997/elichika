package asset

import (
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/handler/common"
	"elichika/locale"
	"elichika/router"
	"elichika/userdata"
	"elichika/utils"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

func getPackUrl(ctx *gin.Context) {
	req := request.GetPackUrlRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	session := ctx.MustGet("session").(*userdata.Session)

	// these are hardcoded links to the original asset version
	cdnMasterVersion := "2d61e7b4e89961c7"
	if ctx.MustGet("locale").(*locale.Locale).MasterVersion == config.MasterVersionJp {
		cdnMasterVersion = "b66ec2295e9a00aa"
	}
	// it's possible to detect TLS using ctx.Request.TLS != nil
	// but that wouldn't work if we're using some external wrapper for TLS instead of elichika itself
	host := *config.Conf.CdnServer
	if host == "elichika" {
		host = "http://" + ctx.Request.Host + "/static"
	} else if host == "elichika_tls" {
		host = "https://" + ctx.Request.Host + "/static"
	}
	resp := response.GetPackUrlResponse{}
	for _, pack := range req.PackNames.Slice {
		assetPack := session.Gamedata.AssetPack[pack]
		if assetPack != nil {
			resp.UrlList.Append(host + "/" + assetPack.MasterVersion + "/" + pack)
		} else {
			resp.UrlList.Append(host + "/" + cdnMasterVersion + "/" + pack)
		}
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/asset/getPackUrl", getPackUrl)
}
