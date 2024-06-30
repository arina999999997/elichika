package asset

// if the server doesn't have any other way of handling range, then we have to do it ourselves
// this assume the server has Accept-Range: bytes, which is true for both elichika or the catfolk cdn.

import (
	"elichika/assetdata"
	"elichika/config"
	"elichika/router"
	"elichika/utils"

	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func staticVirtual(ctx *gin.Context) {
	file := ctx.Param("fileName")
	downloadData := assetdata.GetDownloadData(file)
	if downloadData.IsEntireFile {
		panic("downloading whole file through the virtual endpoint")
	}

	host := *config.Conf.CdnServer
	if host == "elichika" {
		host = "http://" + ctx.Request.Host + "/static"
	} else if host == "elichika_tls" {
		host = "https://" + ctx.Request.Host + "/static"
	}
	client := &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", host, cdnMasterVersionMapping[downloadData.Locale], downloadData.File), nil)
	utils.CheckErr(err)
	request.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", downloadData.Start, downloadData.Start+downloadData.Size-1))
	response, err := client.Do(request)
	utils.CheckErr(err)
	defer response.Body.Close()
	if response.StatusCode != http.StatusPartialContent { // http.StatusPartialContent
		panic("wrong status received")
	}
	body, err := io.ReadAll(response.Body)
	utils.CheckErr(err)

	ctx.Header("Content-Length", fmt.Sprint(downloadData.Size))
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Writer.Write(body)
}

func init() {
	router.AddHandler("/static_virtual", "GET", "/:fileName", staticVirtual)
}
