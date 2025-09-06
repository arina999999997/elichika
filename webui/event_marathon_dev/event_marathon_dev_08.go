//go:build dev

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
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	RuleDescriptionPageAssetPaths = []string{
		"8FA",
		";b;",
	}
)

func EventMarathonDev08GET(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	// pageInt, err := strconv.Atoi(page)
	// utils.CheckErr(err)
	form := image_form.ImageForm{
		FormId:    "rule_description_form",
		DataLabel: "Rule Description Page " + page,
		DataId:    "rule_description_page_" + page,
	}

	for _, assetPath := range image_form.GetAssetPathsFromSamePackage(TopStatus.TitleImagePath.V.Value) {
		form.AddImageByAssetPathFilterBySize("", assetPath, 1184, 520)
	}
	for _, assetPath := range RuleDescriptionPageAssetPaths {
		form.AddImageByAssetPath("", assetPath)
	}

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_marathon_dev.html", gin.H{
		"body": msg,
	})
}

func EventMarathonDev08POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)
	page := ctx.DefaultQuery("page", "1")
	pageInt, err := strconv.Atoi(page)
	utils.CheckErr(err)
	assetPath := form.Value["rule_description_page_"+page][0]
	if assetPath == "" {
		// done
		ctx.Header("Location", "/webui/event_marathon_dev/09")
	} else {
		if len(TopStatus.EventMarathonRuleDescriptionPageMasterRows.Slice) >= pageInt {
			TopStatus.EventMarathonRuleDescriptionPageMasterRows.Slice = TopStatus.EventMarathonRuleDescriptionPageMasterRows.Slice[:pageInt-1]
		}
		TopStatus.EventMarathonRuleDescriptionPageMasterRows.Append(client.EventMarathonRuleDescriptionPageMasterRow{
			Page: int32(pageInt),
			Title: client.LocalizedText{
				DotUnderText: fmt.Sprintf("Event Rules %s/?", page),
			},
			ImageAssetPath: client.TextureStruktur{
				V: generic.NewNullable[string](assetPath),
			},
		})
		ctx.Header("Location", fmt.Sprintf("/webui/event_marathon_dev/08?page=%d", pageInt+1))
	}
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/08", EventMarathonDev08GET)
		router.AddHandler("/webui/event_marathon_dev", "POST", "/08", EventMarathonDev08POST)
	}
}
