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

var (
	RuleDescriptionPageAssetPaths = []string{
		"<:t",
		"`^Q",
		"o:K",
		"!3)",
	}
)

func EventMiningDev07GET(ctx *gin.Context) {
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
	ctx.HTML(http.StatusOK, "event_mining_dev.html", gin.H{
		"body": msg,
	})
}

func EventMiningDev07POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)
	page := ctx.DefaultQuery("page", "1")
	pageInt, err := strconv.Atoi(page)
	utils.CheckErr(err)
	assetPath := form.Value["rule_description_page_"+page][0]
	if assetPath == "" {
		// done
		ctx.Header("Location", "/webui/event_mining_dev/08")
	} else {
		if len(TopStatus.EventMiningRuleDescriptionPageMasterRows.Slice) >= pageInt {
			TopStatus.EventMiningRuleDescriptionPageMasterRows.Slice = TopStatus.EventMiningRuleDescriptionPageMasterRows.Slice[:pageInt-1]
		}
		TopStatus.EventMiningRuleDescriptionPageMasterRows.Append(client.EventMiningRuleDescriptionPageMasterRow{
			Page: int32(pageInt),
			Title: client.LocalizedText{
				DotUnderText: fmt.Sprintf("Event Rules %s/?", page),
			},
			ImageAssetPath: client.TextureStruktur{
				V: generic.NewNullable[string](assetPath),
			},
		})
		ctx.Header("Location", fmt.Sprintf("/webui/event_mining_dev/07?page=%d", pageInt+1))
	}
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/07", EventMiningDev07GET)
		router.AddHandler("/webui/event_mining_dev", "POST", "/07", EventMiningDev07POST)
	}
}
