package event_marathon_dev

import (
	"elichika/config"
	"elichika/locale"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/image_form"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EventMarathonDev09GET(ctx *gin.Context) {
	type EventBoosterItem struct {
		Id                 int32  `xorm:"'id'"`
		ThumbnailAssetPath string `xorm:"'thumbnail_asset_path'"`
	}
	boosterItems := []EventBoosterItem{}
	path := locale.Locales["en"].Path
	err := locale.GetEngine(path + "masterdata.db").Table("m_event_marathon_booster_item").OrderBy("id").Find(&boosterItems)
	utils.CheckErr(err)

	form := image_form.ImageForm{
		FormId:    "booster_item_id_form",
		DataLabel: "Booster item id",
		DataId:    "booster_item_id",
	}

	for _, booster := range boosterItems {
		form.AddImageByAssetPath(fmt.Sprint(booster.Id), booster.ThumbnailAssetPath)
	}

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_marathon_dev.html", gin.H{
		"body": msg,
	})
}

func EventMarathonDev09POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)

	boosterId := form.Value["booster_item_id"][0]
	boosterIdInt, err := strconv.Atoi(boosterId)
	utils.CheckErr(err)
	BoosterItemId = int32(boosterIdInt)
	ctx.Header("Location", "/webui/event_marathon_dev/10")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/09", EventMarathonDev09GET)
		router.AddHandler("/webui/event_marathon_dev", "POST", "/09", EventMarathonDev09POST)
	}
}
