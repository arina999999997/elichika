package event_marathon_dev

// TODO: note sure why but partial viewing of bonus doesn't work, even after adding proper bootstrap information
// maybe there need to be some object actually created somewhere?
// for now let's just raw dog it till the end once we reach this step
import (
	"elichika/client"
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

func EventMarathonDev10GET(ctx *gin.Context) {
	cardId := ctx.DefaultQuery("card", "0")
	cardIdInt, err := strconv.Atoi(cardId)
	utils.CheckErr(err)
	cards := []int32{}
	path := locale.Locales["en"].Path
	err = locale.GetEngine(path+"masterdata.db").Table("m_event_marathon_bonus_card").Where("event_marathon_master_id = ? AND grade = 0", TopStatus.EventId).Cols("card_master_id").OrderBy("value DESC").Find(&cards)
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
	form.AddImageByAssetPath("210", assetPath)
	form.AddImageByAssetPath("210", awakenAssetPath)
	form.AddImageByAssetPath("230", assetPath)
	form.AddImageByAssetPath("230", awakenAssetPath)

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_marathon_dev.html", gin.H{
		"body": msg,
	})
}

func EventMarathonDev10POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)

	cardId := ctx.DefaultQuery("card", "0")
	cardIdInt, err := strconv.Atoi(cardId)
	utils.CheckErr(err)
	cards := []int32{}
	path := locale.Locales["en"].Path
	err = locale.GetEngine(path+"masterdata.db").Table("m_event_marathon_bonus_card").Where("event_marathon_master_id = ? AND grade = 0", TopStatus.EventId).Cols("card_master_id").OrderBy("value DESC").Find(&cards)
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

	if len(TopStatus.EventMarathonBonusPopupOrderCardMaterRows.Slice) > cardIdInt {
		TopStatus.EventMarathonRuleDescriptionPageMasterRows.Slice = TopStatus.EventMarathonRuleDescriptionPageMasterRows.Slice[:cardIdInt]
		TopStatus.EventMarathonBonusPopupOrderMemberMaterRows.Slice = TopStatus.EventMarathonBonusPopupOrderMemberMaterRows.Slice[:cardIdInt]
	}

	TopStatus.EventMarathonBonusPopupOrderCardMaterRows.Append(client.EventMarathonBonusPopupOrderCardMaterRow{
		CardMatserId: cardMasterId,
		DisplayLine:  displayLine,
		DisplayOrder: displayOrder,
		IsGacha:      isGacha,
	})

	TopStatus.EventMarathonBonusPopupOrderMemberMaterRows.Append(client.EventMarathonBonusPopupOrderMemberMaterRow{
		MemberMatserId: (cardMasterId / 10000) % 1000,
		DisplayLine:    3,
		DisplayOrder:   (cardMasterId / 10000) % 1000,
	})
	if cardIdInt+1 == len(cards) {
		ctx.Header("Location", "/webui/event_marathon_dev/11")
	} else {
		ctx.Header("Location", fmt.Sprintf("/webui/event_marathon_dev/10?card=%d", cardIdInt+1))
	}
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/10", EventMarathonDev10GET)
		router.AddHandler("/webui/event_marathon_dev", "POST", "/10", EventMarathonDev10POST)
	}
}
