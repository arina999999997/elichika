package event_marathon_dev

import (
	"elichika/config"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/image_form"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	EventNameLanguages = [4]string{"ja", "en", "ko", "zh"}
)

func EventMarathonDev02GET(ctx *gin.Context) {
	language := ctx.Request.URL.RawQuery
	if language == "" {
		language = "ja"
	}
	languageId := 0
	for ; (languageId < 4) && (EventNameLanguages[languageId] != language); languageId++ {
	}
	if languageId == 4 {
		languageId = 0
	}

	form := image_form.ImageForm{
		FormId:    "event_name_" + EventNameLanguages[languageId] + "_form",
		DataLabel: "Event name (" + EventNameLanguages[languageId] + ")",
		DataId:    "event_name_" + EventNameLanguages[languageId],
	}
	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_marathon_dev.html", gin.H{
		"body": msg,
	})
}

func EventMarathonDev02POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	language := ctx.Request.URL.RawQuery
	if language == "" {
		language = "ja"
	}
	languageId := 0
	for ; (languageId < 4) && (EventNameLanguages[languageId] != language); languageId++ {
	}
	if languageId == 4 {
		languageId = 0
	}
	fmt.Println(form.Value)
	EventName[EventNameLanguages[languageId]] = form.Value["event_name_"+EventNameLanguages[languageId]][0]
	utils.CheckErr(err)
	if languageId == 3 {
		ctx.Header("Location", "/webui/event_marathon_dev/03")
	} else {
		ctx.Header("Location", "/webui/event_marathon_dev/02?"+EventNameLanguages[languageId+1])
	}
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/02", EventMarathonDev02GET)
		router.AddHandler("/webui/event_marathon_dev", "POST", "/02", EventMarathonDev02POST)
	}
}
