//go:build dev

package event_mining_dev

import (
	"elichika/client"
	"elichika/config"
	"elichika/locale"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/form/image_form"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EventMiningDev08GET(ctx *gin.Context) {
	cardId := ctx.DefaultQuery("card", "0")
	cardIdInt, err := strconv.Atoi(cardId)
	utils.CheckErr(err)
	cards := []int32{}
	path := locale.Locales["en"].Path
	err = locale.GetEngine(path+"masterdata.db").Table("m_event_mining_bonus_card").Where("event_mining_master_id = ? AND grade = 0", TopStatus.EventId).Cols("card_master_id").OrderBy("value DESC").Find(&cards)
	utils.CheckErr(err)
	if cardIdInt >= len(cards) {
		panic("already done")
	}
	cardMasterId := cards[cardIdInt]
	form := image_form.ImageForm{
		FormId:    "card_bonus_popup_order_form",
		DataLabel: fmt.Sprintf("Card Bonus Popup Order for %d", cardMasterId),
		DataId:    "card_bonus_popup_order_" + cardId,
	}
	assetPath, awakenAssetPath := image_form.GetCardAssetPaths(cardMasterId)

	// DisplayLine DisplayOrder IsGacha
	form.AddImageByAssetPath("111", assetPath)
	form.AddImageByAssetPath("111", awakenAssetPath)
	form.AddImageByAssetPath("121", assetPath)
	form.AddImageByAssetPath("121", awakenAssetPath)
	form.AddLinebreak()
	form.AddImageByAssetPath("231", assetPath)
	form.AddImageByAssetPath("231", awakenAssetPath)
	form.AddLinebreak()
	form.AddImageByAssetPath("110", assetPath)
	form.AddImageByAssetPath("110", awakenAssetPath)
	form.AddLinebreak()
	form.AddImageByAssetPath("220", assetPath)
	form.AddImageByAssetPath("220", awakenAssetPath)
	form.AddImageByAssetPath("230", assetPath)
	form.AddImageByAssetPath("230", awakenAssetPath)

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_mining_dev.html", gin.H{
		"body": msg,
	})
}

func EventMiningDev08POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)

	cardId := ctx.DefaultQuery("card", "0")
	cardIdInt, err := strconv.Atoi(cardId)
	utils.CheckErr(err)
	cards := []int32{}
	path := locale.Locales["en"].Path
	err = locale.GetEngine(path+"masterdata.db").Table("m_event_mining_bonus_card").Where("event_mining_master_id = ? AND grade = 0", TopStatus.EventId).Cols("card_master_id").OrderBy("value DESC").Find(&cards)
	utils.CheckErr(err)
	if cardIdInt >= len(cards) {
		panic("already done")
	}
	cardMasterId := cards[cardIdInt]
	positionType := form.Value["card_bonus_popup_order_"+cardId][0]
	positionTypeInt, err := strconv.Atoi(positionType)
	displayLine := int32(positionTypeInt / 100)
	displayOrder := int32(positionTypeInt/10) % 10
	isGacha := (positionTypeInt % 10) == 1

	if len(TopStatus.EventMiningBonusPopupOrderCardMaterRows.Slice) > cardIdInt {
		TopStatus.EventMiningBonusPopupOrderCardMaterRows.Slice = TopStatus.EventMiningBonusPopupOrderCardMaterRows.Slice[:cardIdInt]
	}

	TopStatus.EventMiningBonusPopupOrderCardMaterRows.Append(client.EventMiningBonusPopupOrderCardMaterRow{
		CardMatserId: cardMasterId,
		DisplayLine:  displayLine,
		DisplayOrder: displayOrder,
		IsGacha:      isGacha,
	})

	if cardIdInt+1 == len(cards) {
		ctx.Header("Location", "/webui/event_mining_dev/09")
	} else {
		ctx.Header("Location", fmt.Sprintf("/webui/event_mining_dev/08?card=%d", cardIdInt+1))
	}
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/08", EventMiningDev08GET)
		router.AddHandler("/webui/event_mining_dev", "POST", "/08", EventMiningDev08POST)
	}
}
