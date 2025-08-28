package event_marathon_dev

import (
	"elichika/client"
	"elichika/config"
	"elichika/generic"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/image_form"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	BoardDecoImageAssetPaths = []string{
		`br{`,
	}
)

func EventMarathonDev06GET(ctx *gin.Context) {

	form := image_form.ImageForm{
		FormId:    "board_deco_image_path_form",
		DataLabel: "Event board image",
		DataId:    "board_deco_image_path",
	}

	for _, assetPath := range BoardDecoImageAssetPaths {
		form.AddImageByAssetPath("", assetPath)
	}

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_marathon_dev.html", gin.H{
		"body": msg,
	})
}

func EventMarathonDev06POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)
	decoImagePath := form.Value["board_deco_image_path"][0]
	if decoImagePath == "" || decoImagePath == "null" {
		TopStatus.BoardStatus.BoardDecoImagePath = client.TextureStruktur{
			V: generic.Nullable[string]{},
		}
	} else {
		TopStatus.BoardStatus.BoardDecoImagePath = client.TextureStruktur{
			V: generic.NewNullable[string](decoImagePath),
		}
	}
	ctx.Header("Location", "/webui/event_marathon_dev/07")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/06", EventMarathonDev06GET)
		router.AddHandler("/webui/event_marathon_dev", "POST", "/06", EventMarathonDev06POST)
	}
}
