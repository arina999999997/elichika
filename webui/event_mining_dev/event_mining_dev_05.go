//go:build dev

package event_mining_dev

import (
	"elichika/client"
	"elichika/config"
	"elichika/generic"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/form/image_form"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// these are the arts for each idol
// aside from that, we might have luck getting the idol list from the package:
// - the package usually contain more than necessary amount of extra idol, but sometime they don't contain any at all
// - so we will pick from the one in the same packs first, and then we will pick from the ont listed here
var topStillCellAssetPaths = []string{
	`!+q`, // 48 honoka
	"5`B", // 45 umi
	`{gT`, // 49 nozomi
	`w~>`, // 48 maki
	`>vs`, // 48 eli
	"`()", // 46 rin
	`6Zx`, // 43 hanayo
	`Csy`, // 39 nico
	`x[7`, // 41 kotori
	`:5n`, // 54 chika
	`<c}`, // 53 ruby
	`bV8`, // 41 kanan
	`;lL`, // 41 hanamaru
	`K/\`, // 43 yoshiko
	`a#;`, // 48 riko
	`;iG`, // 47 dia
	`b[l`, // 37 you
	`w<M`, // 49 mari
	"DO`", // 52 karin
	`6j\`, // 50 emma
	`QI4`, // 49 shizuku
	`;'V`, // 49 ai
	`wYY`, // 49 kasumi
	`%lC`, // 43 ayumu
	`VqW`, // 41 rina
	`Y5u`, // 40 setsuna
	`E&W`, // 34 kanata
	`s3=`, // 15 shioriko
}

func EventMiningDev05GET(ctx *gin.Context) {
	cellId := ctx.DefaultQuery("top_still_sub_cell_id", "1")

	form := image_form.ImageForm{
		FormId:    "event_top_still_sub_cell_form",
		DataLabel: "Event top still cells (note: select by appearance order AT THE END, top to bottom)",
		DataId:    "event_top_still_sub_cell_" + fmt.Sprint(cellId),
	}
	// these id are not consistent, so we just go from 1 to 15
	// id = 1 and id = 2 is often used by the event itself
	assetPaths := image_form.GetAssetPathsFromSamePackage(TopStatus.TitleImagePath.V.Value)
	alreadyUsed := map[string]bool{}
	for _, row := range TopStatus.EventMiningTopStillCellMasterRows.Slice {
		alreadyUsed[row.ImageThumbnailAssetPath.V.Value] = true
	}
	for _, row := range TopStatus.EventMiningTopStillSubCellMasterRows.Slice {
		alreadyUsed[row.ImageThumbnailAssetPath.V.Value] = true
	}
	for _, assetPath := range assetPaths {
		if alreadyUsed[assetPath] {
			continue
		}
		form.AddImageByAssetPathFilterBySize(assetPath, assetPath, 532, 532)
	}

	for _, assetPath := range topStillCellAssetPaths {
		if alreadyUsed[assetPath] {
			continue
		}
		form.AddImageByAssetPath(assetPath, assetPath)
	}

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_mining_dev.html", gin.H{
		"body": msg,
	})
}

func EventMiningDev05POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)
	cellId := ctx.DefaultQuery("top_still_sub_cell_id", "1")
	cellIdInt, err := strconv.Atoi(cellId)
	utils.CheckErr(err)
	assetPath := form.Value["event_top_still_sub_cell_"+cellId][0]
	if cellIdInt == 1 {
		TopStatus.EventMiningTopStillSubCellMasterRows.Slice = nil
	}
	priority := 16 - cellIdInt
	if cellIdInt <= 2 {
		if cellIdInt == 1 {
			priority = 30
		} else {
			priority = 20
		}
	}
	TopStatus.EventMiningTopStillSubCellMasterRows.Append(
		client.EventMiningTopStillSubCellMasterRow{
			EventMiningMasterId: TopStatus.EventId,
			PanelSetId:          2,
			ThumbnailCellId:     TopStatus.EventId*10000 + 1100 + int32(cellIdInt),
			Priority:            int32(priority),
			ImageThumbnailAssetPath: client.TextureStruktur{
				V: generic.NewNullable(assetPath),
			},
			IsPopup: true,
			PopupImageThumbnailAssetPath: client.TextureStruktur{
				V: generic.NewNullable(assetPath),
			},
		})
	if cellIdInt == 15 { // done
		ctx.Header("Location", "/webui/event_mining_dev/06")
	} else {
		ctx.Header("Location", fmt.Sprintf("/webui/event_mining_dev/05?top_still_sub_cell_id=%d", cellIdInt+1))
	}
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/05", EventMiningDev05GET)
		router.AddHandler("/webui/event_mining_dev", "POST", "/05", EventMiningDev05POST)
	}
}
