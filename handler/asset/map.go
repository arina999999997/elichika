package asset

import (
	"elichika/assetdata"
	"elichika/router"

	"fmt"

	"github.com/gin-gonic/gin"
)

// acting as the cdn, we need a map from file to actual files
func staticMap(ctx *gin.Context) {
	file := ctx.Param("fileName")
	downloadData := assetdata.GetDownloadData(file)
	if downloadData.IsEntireFile {
		panic("entire file downloaded through map endpoint")
	}

	sendRange(ctx, fmt.Sprintf("static/%s/%s", cdnMasterVersionMapping[downloadData.Locale], downloadData.File), downloadData.Start, downloadData.Size)
}

func init() {
	router.AddHandler("/static_map", "GET", "/:fileName", staticMap)
}
