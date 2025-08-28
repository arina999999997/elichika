package event_marathon_dev

import (
	"elichika/client"
	"elichika/config"
	"elichika/generic"
	"elichika/router"
	"elichika/utils"
	// "elichika/locale"
	"elichika/webui/image_form"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	BoardBaseImageAssetPaths = []string{
		`+6I`,
		// `Fca`, this is similar to +6I but it's bigger for some reason
		`')k`,
		`r0|`,
		`PFW`,
		`pQ\`,
		`w'y`,
		`')k`,
		`wT*`,
	}
)

func EventMarathonDev05GET(ctx *gin.Context) {

	form := image_form.ImageForm{
		FormId:    "board_base_image_path_form",
		DataLabel: "Event board image",
		DataId:    "board_base_image_path",
	}

	// TODO: fetch from same package if necessary
	// scenariors := []string{}
	// path := locale.Locales["en"].Path
	// err := locale.GetEngine(path + "masterdata.db").Table("m_story_event_history_detail").Where("event_master_id = ?", TopStatus.EventId).Cols("scenario_script_asset_path").Find(&scenariors)
	// utils.CheckErr(err)
	// assetDb := locale.GetEngine(path + "asset_a_en.db")
	// textures := []string{}
	// texturesMap := map[string]bool{}
	// for _, script := range scenariors {
	// 	err = assetDb.Table("adv_graphic").Where("script_name = ?", script).Cols("resource").Find(&textures)
	// 	utils.CheckErr(err)
	// 	for _, assetPath := range textures {
	// 		texturesMap[assetPath] = true
	// 	}
	// }
	for _, assetPath := range BoardBaseImageAssetPaths {
		form.AddImageByAssetPath("", assetPath)
	}

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_marathon_dev.html", gin.H{
		"body": msg,
	})
}

func EventMarathonDev05POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)
	TopStatus.BoardStatus.BoardBaseImagePath = client.TextureStruktur{
		V: generic.NewNullable[string](form.Value["board_base_image_path"][0]),
	}
	ctx.Header("Location", "/webui/event_marathon_dev/06")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/05", EventMarathonDev05GET)
		router.AddHandler("/webui/event_marathon_dev", "POST", "/05", EventMarathonDev05POST)
	}
}
