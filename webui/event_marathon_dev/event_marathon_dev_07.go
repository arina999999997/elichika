//go:build dev

package event_marathon_dev

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/generic"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/image_form"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	BoardThingTypeById = []int32{
		enum.EventMarathonBoardPositionTypeMemo,
		enum.EventMarathonBoardPositionTypeMemo,
		enum.EventMarathonBoardPositionTypeMemo,
		enum.EventMarathonBoardPositionTypePicture,
		enum.EventMarathonBoardPositionTypePicture,
		enum.EventMarathonBoardPositionTypePicture,
		enum.EventMarathonBoardPositionTypePicture,
		enum.EventMarathonBoardPositionTypePicture,
		enum.EventMarathonBoardPositionTypePicture,
		enum.EventMarathonBoardPositionTypePicture,
	}
	BoardThingCount = len(BoardThingTypeById)
)

func EventMarathonDev07GET(ctx *gin.Context) {
	boardThingId := ctx.DefaultQuery("board_thing_id", "0")
	boardThingIdInt, err := strconv.Atoi(boardThingId)
	utils.CheckErr(err)
	phase := ctx.DefaultQuery("phase", "image")
	label := "Event board "
	if BoardThingTypeById[boardThingIdInt] == enum.EventMarathonBoardPositionTypeMemo {
		label += "memo "
	} else {
		label += "picture "
	}
	form := image_form.ImageForm{
		FormId:    "board_thing_form",
		DataLabel: label + phase,
		DataId:    "board_thing_" + boardThingId + "_" + phase,
	}

	if phase == "image" {
		for _, assetPath := range image_form.GetAssetPathsFromSamePackage(TopStatus.TitleImagePath.V.Value) {
			if BoardThingTypeById[boardThingIdInt] == enum.EventMarathonBoardPositionTypeMemo {
				form.AddImageByAssetPathFilterBySize("", assetPath, 200, 200)
			} else {
				form.AddImageByAssetPathFilterBySize("", assetPath, 256, 160)
			}
		}
	} else {
		assetPath := TopStatus.BoardStatus.BoardThingMasterRows.Slice[boardThingIdInt].ImageThumbnailAssetPath.V.Value
		if BoardThingTypeById[boardThingIdInt] == enum.EventMarathonBoardPositionTypeMemo {
			for i := 1; i <= 9; i++ {
				form.AddImageByAssetPath(fmt.Sprintf("%d", i), assetPath)
				if i == 3 || i == 6 {
					form.AddLinebreak()
				}
			}
		} else {
			for i := 1; i <= 7; i++ {
				form.AddImageByAssetPath(fmt.Sprintf("%d", i), assetPath)
				if i == 3 || i == 5 {
					form.AddLinebreak()
				}
			}
		}
	}

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_marathon_dev.html", gin.H{
		"body": msg,
	})
}

func EventMarathonDev07POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)
	boardThingId := ctx.DefaultQuery("board_thing_id", "0")
	boardThingIdInt, err := strconv.Atoi(boardThingId)
	utils.CheckErr(err)
	phase := ctx.DefaultQuery("phase", "image")
	if phase == "image" {
		assetPath := form.Value["board_thing_"+boardThingId+"_"+phase][0]
		if len(TopStatus.BoardStatus.BoardThingMasterRows.Slice) > boardThingIdInt {
			TopStatus.BoardStatus.BoardThingMasterRows.Slice = TopStatus.BoardStatus.BoardThingMasterRows.Slice[:boardThingIdInt]
		}
		TopStatus.BoardStatus.BoardThingMasterRows.Append(client.EventMarathonBoardMemorialThingsMasterRow{
			EventMarathonBoardPositionType: BoardThingTypeById[boardThingIdInt],
			Position:                       1,
			Priority:                       int32(boardThingIdInt + 1),
			ImageThumbnailAssetPath: client.TextureStruktur{
				V: generic.NewNullable[string](assetPath),
			},
		})
		ctx.Header("Location", fmt.Sprintf("/webui/event_marathon_dev/07?board_thing_id=%s&phase=position", boardThingId))
	} else {
		position, err := strconv.Atoi(form.Value["board_thing_"+boardThingId+"_"+phase][0])
		utils.CheckErr(err)
		TopStatus.BoardStatus.BoardThingMasterRows.Slice[boardThingIdInt].Position = int32(position)
		if boardThingIdInt+1 == BoardThingCount {
			ctx.Header("Location", "/webui/event_marathon_dev/08")
		} else {
			ctx.Header("Location", fmt.Sprintf("/webui/event_marathon_dev/07?board_thing_id=%d&phase=image", boardThingIdInt+1))
		}
	}
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/07", EventMarathonDev07GET)
		router.AddHandler("/webui/event_marathon_dev", "POST", "/07", EventMarathonDev07POST)
	}
}
