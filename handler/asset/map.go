package asset

import (
	"elichika/assetdata"
	"elichika/config"
	"elichika/log"
	"elichika/router"

	"fmt"

	"github.com/gin-gonic/gin"
)

// acting as the cdn, we need a map from file to actual files
func staticMap(ctx *gin.Context) {
	if (*config.Conf.CdnServer != "elichika") && (*config.Conf.CdnServer != "elichika_tls") {
		log.Panic("staticMap is not allowed because CDN is not elichika or elichika tls")
	}
	file := ctx.Param("fileName")
	downloadData := assetdata.GetDownloadData(file)
	if downloadData.IsEntireFile {
		log.Panic("entire file downloaded through map endpoint")
	}

	sendRange(ctx, fmt.Sprintf("static/%s", downloadData.File), downloadData.Start, downloadData.Size)
}

func init() {
	router.AddHandler("/static_map", "GET", "/:fileName", staticMap)
}
