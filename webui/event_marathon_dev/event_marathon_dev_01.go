//go:build dev

package event_marathon_dev

import (
	"elichika/config"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/image_form"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EventMarathonDev01GET(ctx *gin.Context) {
	form := image_form.ImageForm{
		FormId:    "event_id_form",
		DataLabel: "Event Id",
		DataId:    "event_id",
	}
	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_marathon_dev.html", gin.H{
		"body": msg,
	})
}

func EventMarathonDev01POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)
	eventId, err := strconv.Atoi(form.Value["event_id"][0])
	utils.CheckErr(err)
	TopStatus.EventId = int32((eventId))
	ctx.Header("Location", "/webui/event_marathon_dev/02")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/", EventMarathonDev01GET)
		router.AddHandler("/webui/event_marathon_dev", "POST", "/", EventMarathonDev01POST)
	}
}
