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
	"strings"

	"github.com/gin-gonic/gin"
)

var addStoryNumberByCellId = map[int]int32{}

func init() {
	addStoryNumberByCellId[16] = 0
	addStoryNumberByCellId[17] = 0
	addStoryNumberByCellId[18] = 1
	addStoryNumberByCellId[19] = 2
	addStoryNumberByCellId[20] = 3
	addStoryNumberByCellId[21] = 7
	addStoryNumberByCellId[22] = 4
	addStoryNumberByCellId[23] = 5
	addStoryNumberByCellId[24] = 6
	addStoryNumberByCellId[25] = 7
	addStoryNumberByCellId[26] = 7
}

func EventMiningDev04GET(ctx *gin.Context) {
	cellId := ctx.DefaultQuery("top_still_cell_id", "16")

	form := image_form.ImageForm{
		FormId:    "event_top_still_cell_form",
		DataLabel: "Event top still cells (note: select by appearance order AT THE END, from bottom up, prefix with animation_ to use animation instead. If using animation, use _animation_split_ to split between the normal and popup assets, and put the event card animation in the middle spot)",
		DataId:    "event_top_still_cell_" + fmt.Sprint(cellId),
	}

	assetPaths := image_form.GetAssetPathsFromSamePackage(TopStatus.TitleImagePath.V.Value)
	alreadyUsed := map[string]bool{}
	for _, row := range TopStatus.EventMiningTopStillCellMasterRows.Slice {
		alreadyUsed[row.ImageThumbnailAssetPath.V.Value] = true
	}
	for _, assetPath := range assetPaths {
		if alreadyUsed[assetPath] {
			continue
		}
		form.AddImageByAssetPathFilterBySize(assetPath, assetPath, 532, 532)
	}

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_mining_dev.html", gin.H{
		"body": msg,
	})
}

func EventMiningDev04POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)
	cellId := ctx.DefaultQuery("top_still_cell_id", "16")
	cellIdInt, err := strconv.Atoi(cellId)
	utils.CheckErr(err)
	assetPath := form.Value["event_top_still_cell_"+cellId][0]
	if cellIdInt == 16 {
		TopStatus.EventMiningTopStillCellMasterRows.Slice = nil
	}
	if !strings.HasPrefix(assetPath, "animation_") {
		TopStatus.EventMiningTopStillCellMasterRows.Append(
			client.EventMiningTopStillCellMasterRow{
				EventMiningMasterId: TopStatus.EventId,
				ThumbnailCellId:     TopStatus.EventId*100 + int32(cellIdInt),
				AddStoryNumber:      addStoryNumberByCellId[cellIdInt],
				Priority:            int32(cellIdInt * 10),
				ImageThumbnailAssetPath: client.TextureStruktur{
					V: generic.NewNullable(assetPath),
				},
				IsPopup: true,
				PopupImageThumbnailAssetPath: client.TextureStruktur{
					V: generic.NewNullable(assetPath),
				},
				IsBig: (cellIdInt == 16) || (cellIdInt == 21) || (cellIdInt == 26),
			})
	} else {
		assetPaths := strings.Split(assetPath[10:], "_animation_split_")
		if len(assetPaths) != 2 {
			panic("must have exactly the normal and pop up asset path")
		}
		TopStatus.EventMiningTopAnimationCellMasterRows.Append(
			client.EventMiningTopAnimationCellMasterRow{
				EventMiningMasterId: TopStatus.EventId,
				ThumbnailCellId:     TopStatus.EventId*100 + int32(cellIdInt),
				// this is actually set to 0 by the client anyway, but server side want it to reveal the cells based on story read
				AddStoryNumber: addStoryNumberByCellId[cellIdInt],
				Priority:       int32(cellIdInt * 10),
				MovieAssetPath: client.MovieStruktur{
					V: assetPaths[0],
				},
				IsPopup: true,
				PopupMovieAssetPath: client.MovieStruktur{
					V: assetPaths[1],
				},
				IsBig: true,
				// the other fields are not necessary
			})
	}
	if cellIdInt == 26 { // done
		ctx.Header("Location", "/webui/event_mining_dev/05")
	} else {
		ctx.Header("Location", fmt.Sprintf("/webui/event_mining_dev/04?top_still_cell_id=%d", cellIdInt+1))
	}
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/04", EventMiningDev04GET)
		router.AddHandler("/webui/event_mining_dev", "POST", "/04", EventMiningDev04POST)
	}
}
