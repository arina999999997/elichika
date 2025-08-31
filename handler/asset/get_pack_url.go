package asset

import (
	"elichika/assetdata"
	"elichika/client/request"
	"elichika/client/response"
	"elichika/config"
	"elichika/handler/common"
	"elichika/log"
	"elichika/router"
	"elichika/utils"

	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func getPackUrl(ctx *gin.Context) {
	req := request.GetPackUrlRequest{}
	err := json.Unmarshal(*ctx.MustGet("reqBody").(*json.RawMessage), &req)
	utils.CheckErr(err)

	host := *config.Conf.CdnServer

	if (host == "elichika") || (host == "elichika_tls") {
		actualHost := ctx.Request.Host
		// ctx.Request.Proto is not what we want, it is HTTP/1.0 and similar and doesn't indicate whether the connection is TLS or not
		actualProto := "http"
		if ctx.Request.TLS != nil {
			actualProto = "https"
		}

		// if the connection is forwarded, we need to return the forwarded host instead
		forwardedHost, hostExists := ctx.Request.Header["X-Forwarded-Host"]
		forwardedProto, protoExists := ctx.Request.Header["X-Forwarded-Proto"]
		if hostExists && len(forwardedHost) > 0 {
			actualHost = forwardedHost[0]
		}
		if protoExists && len(forwardedProto) > 0 {
			actualProto = forwardedProto[0]
		}
		host = actualProto + "://" + actualHost + "/static"
	}
	resp := response.GetPackUrlResponse{}
	for _, pack := range req.PackNames.Slice {
		downloadData := assetdata.GetDownloadData(pack)
		if downloadData.IsEntireFile { // always use the cdn server as is for all whole files
			resp.UrlList.Append(fmt.Sprintf("%s/%s", host, pack))
			continue
		}
		if *config.Conf.CdnPartialFileCapability == "static_file" {
			// if the cdn has static partial files then just give a normal request
			// this is simple but require more storage on the cdn server
			resp.UrlList.Append(fmt.Sprintf("%s/%s", host, pack))
		} else if *config.Conf.CdnPartialFileCapability == "mapped_file" {
			// end point is /static_map/<file>
			// if the cdn has mapping from partial files, (i.e. elichika itself) then just send the file name to this mapped api
			// having a separate endpoint help with some server impl.
			// if the server can use one endpoint for both normal and partial files, then using "static_file" should have the same effect.
			// this will require the cdn server to have some sort of mapping on hand
			// but it will also allow the cdn server to do some caching, as the urls are the same
			resp.UrlList.Append(fmt.Sprintf("%s_map/%s", host, pack))
		} else if *config.Conf.CdnPartialFileCapability == "has_range_api" {
			// end point is /static_api?&file=<file>&start=<start>&size=<size>
			// this allow the cdn server to implement a simple range download function.
			// it can be cached too if, but it'll be more vulnerable to random queries that doesn't represent an actual file.
			resp.UrlList.Append(fmt.Sprintf("%s_api?file=%s&start=%d&size=%d", host,
				downloadData.File, downloadData.Start, downloadData.Size))
		} else if *config.Conf.CdnPartialFileCapability == "nothing" {
			// the cdn server can't deal with partial files, so it's up to elichika to help it
			// TODO(extra): this assume the server is http or it can auto upgrade to https if necessary
			// i.e. this address will be served correctly
			virtualHost := "http://" + ctx.Request.Host + "/static"
			resp.UrlList.Append(fmt.Sprintf("%s_virtual/%s", virtualHost, pack))
		} else {
			log.Panic("wrong cdn_partial_file_capability")
		}
	}

	common.JsonResponse(ctx, &resp)
}

func init() {
	router.AddHandler("/", "POST", "/asset/getPackUrl", getPackUrl)
}
