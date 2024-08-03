package asset

import (
	"elichika/assetdata"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/handler/common"
	"elichika/router"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

var cdnMasterVersionMapping = map[string]string{}

func init() {
	cdnMasterVersionMapping["en"] = "2d61e7b4e89961c7"
	cdnMasterVersionMapping["ko"] = "2d61e7b4e89961c7"
	cdnMasterVersionMapping["zh"] = "2d61e7b4e89961c7"

	// // TODO(cdn): make this change globally
	// cdnMasterVersionMapping["en"] = "b66ec2295e9a00aa"
	// cdnMasterVersionMapping["ko"] = "b66ec2295e9a00aa"
	// cdnMasterVersionMapping["zh"] = "b66ec2295e9a00aa"
	cdnMasterVersionMapping["ja"] = "b66ec2295e9a00aa"
}

func getPackUrl(ctx *gin.Context) {
	req := request.GetPackUrlRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

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
		downloadData := assetdata.GetDownloadData(pack)
		if downloadData.IsEntireFile { // always use the cdn server as is for all whole files
			resp.UrlList.Append(fmt.Sprintf("%s/%s/%s", host, cdnMasterVersionMapping[downloadData.Locale], pack))
			continue
		}
		if *config.Conf.CdnPartialFileCapability == "static_file" {
			// if the cdn has static partial files then just give a normal request
			// this is simple but require more storage on the cdn server
			resp.UrlList.Append(fmt.Sprintf("%s/%s/%s", host, cdnMasterVersionMapping[downloadData.Locale], pack))
		} else if *config.Conf.CdnPartialFileCapability == "mapped_file" {
			// end point is /static_map/<file>
			// if the cdn has mapping from partial files, (i.e. elichika itself) then just send the file name to this mapped api
			// having a separate endpoint help with some server impl.
			// if the server can use one endpoint for both normal and partial files, then using "static_file" should have the same effect.
			// this will require the cdn server to have some sort of mapping on hand
			// but it will also allow the cdn server to do some caching, as the urls are the same
			resp.UrlList.Append(fmt.Sprintf("%s_map/%s", host, pack))
		} else if *config.Conf.CdnPartialFileCapability == "has_range_api" {
			// end point is /static_api?master=<master_version>&file=<file>&start=<start>&size=<size>
			// this allow the cdn server to implement a simple range download function.
			// it can be cached too if, but it'll be more vulnerable to random queries that doesn't represent an actual file.
			resp.UrlList.Append(fmt.Sprintf("%s_api?master=%s&file=%s&start=%d&size=%d", host, cdnMasterVersionMapping[downloadData.Locale],
				downloadData.File, downloadData.Start, downloadData.Size))
		} else if *config.Conf.CdnPartialFileCapability == "nothing" {
			// the cdn server can't deal with partial files, so it's up to elichika to help it
			// TODO(extra): this assume the server is http or it can auto upgrade to https if necessary
			// i.e. this address will be served correctly
			virtualHost := "http://" + ctx.Request.Host + "/static"
			resp.UrlList.Append(fmt.Sprintf("%s_virtual/%s", virtualHost, pack))

		} else {
			panic("wrong cdn_partial_file_capability")
		}
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/asset/getPackUrl", getPackUrl)
}
