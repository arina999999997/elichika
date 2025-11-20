//go:build dev

package event_mining_dev

import (
	"elichika/client"
	"elichika/config"
	"elichika/generic"
	"elichika/locale"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/form/image_form"

	"fmt"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func EventMiningDev03GET(ctx *gin.Context) {

	form := image_form.ImageForm{
		FormId:    "background_image_path_form",
		DataLabel: "Event background image",
		DataId:    "background_image_path",
	}
	scenariors := []string{}
	path := locale.Locales["en"].Path
	err := locale.GetEngine(path+"masterdata.db").Table("m_story_event_history_detail").Where("event_master_id = ?", TopStatus.EventId).Cols("scenario_script_asset_path").Find(&scenariors)
	utils.CheckErr(err)
	assetDb := locale.GetEngine(path + "asset_a_en.db")
	textures := []string{}
	texturesMap := map[string]bool{}
	for _, script := range scenariors {
		err = assetDb.Table("adv_graphic").Where("script_name = ?", script).Cols("resource").Find(&textures)
		utils.CheckErr(err)
		for _, assetPath := range textures {
			texturesMap[assetPath] = true
		}
	}
	textures = nil
	for assetPath := range texturesMap {
		textures = append(textures, assetPath)
	}
	sort.Strings(textures)
	for _, assetPath := range textures {
		form.AddImageByAssetPathFilterBySize("", assetPath, 1800, 900)
	}

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_mining_dev.html", gin.H{
		"body": msg,
	})
}

func EventMiningDev03POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)
	TopStatus.BackgroundImagePath = client.TextureStruktur{
		V: generic.NewNullable[string](form.Value["background_image_path"][0]),
	}
	ctx.Header("Location", "/webui/event_mining_dev/04")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/03", EventMiningDev03GET)
		router.AddHandler("/webui/event_mining_dev", "POST", "/03", EventMiningDev03POST)
	}
}
