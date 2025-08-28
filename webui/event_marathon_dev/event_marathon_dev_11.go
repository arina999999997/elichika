package event_marathon_dev

import (
	"elichika/config"
	"elichika/enum"
	"elichika/locale"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/image_form"

	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	PointRewardLines = []string{}
)

// This is where we build the reward for the event
// TODO: for now only support early moegirl, probably should select the method first if we have method
func EventMarathonDev11GET(ctx *gin.Context) {

	lineId := ctx.DefaultQuery("line_id", "-1")
	lineIdInt, err := strconv.Atoi(lineId)
	utils.CheckErr(err)
	msg := ""
	if lineIdInt == -1 {
		// no item id yet, we will parse it
		msg = `<form id="point_reward_form" method="POST" enctype="multipart/form-data">`
		msg += "\n"
		msg += `<div><label>Upload moegirl old format file: <input type="file" name="file"/> <input type="text" name="point_reward_data"/>`
		msg += "\n"
	} else {
		// we parsed some of the data already, but this item can't be parsed automatically, so we have to resolve it manually
		form := image_form.ImageForm{
			FormId:    "point_reward_form",
			DataLabel: fmt.Sprintf("Item for %s", MoegirlDataLines[lineIdInt]),
			DataId:    "point_reward_data",
		}
		// TODO: add the background image for the event too

		// usually, this is just the BoosterItemId or the reward cards, so we can pull from that
		// booster item id
		path := locale.Locales["en"].Path
		var boosterItemAssetPath string
		exists, err := locale.GetEngine(path+"masterdata.db").Table("m_event_marathon_booster_item").Where("id = ?", BoosterItemId).Cols("thumbnail_asset_path").Get(&boosterItemAssetPath)
		utils.CheckErr(err)
		if exists {
			form.AddImageByAssetPath(fmt.Sprintf("%d,%d", enum.ContentTypeEventMarathonBooster, BoosterItemId), boosterItemAssetPath)
		}
		// background
		backgroundId := 0
		exists, err = locale.GetEngine(path+"masterdata.db").Table("m_background").Where("background_asset_path = ?", TopStatus.BackgroundImagePath.V.Value).Cols("id").Get(&backgroundId)
		utils.CheckErr(err)
		if exists {
			assetPath := ""
			locale.GetEngine(path+"masterdata.db").Table("m_custom_background").Where("background_id = ?", backgroundId).Cols("thumbnail_asset_path").Get(&assetPath)
			form.AddImageByAssetPath(fmt.Sprintf("%d,%d", enum.ContentTypeCustomBackground, backgroundId), assetPath)
		}
		// card
		for _, item := range TopStatus.EventMarathonBonusPopupOrderCardMaterRows.Slice {
			if item.IsGacha {
				continue
			}
			assetPath, awakenAssetPath := image_form.GetCardAssetPaths(item.CardMatserId)
			form.AddImageByAssetPath(fmt.Sprintf("%d,%d", enum.ContentTypeCard, item.CardMatserId), assetPath)
			form.AddImageByAssetPath(fmt.Sprintf("%d,%d", enum.ContentTypeCard, item.CardMatserId), awakenAssetPath)
		}
		msg = form.GetHTML()
	}

	ctx.Header("Content-Type", "text/html")
	ctx.HTML(http.StatusOK, "event_marathon_dev.html", gin.H{
		"body": msg,
	})
}

func EventMarathonDev11POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	lineId := ctx.DefaultQuery("line_id", "-1")
	lineIdInt, err := strconv.Atoi(lineId)
	utils.CheckErr(err)
	data := form.Value["point_reward_data"][0]
	fmt.Println(form)
	if lineIdInt == -1 {
		fileHeader := form.File["file"][0]
		file, err := fileHeader.Open()
		utils.CheckErr(err)
		bytes := make([]byte, fileHeader.Size)
		_, err = file.Read(bytes)
		utils.CheckErr(err)
		lines := strings.Split(string(bytes), "\r\n")
		MoegirlDataLines = []string{}
		PointRewardLines = []string{}
		for _, line := range lines {
			reward := MoegirlGetPointReward(line)
			if reward == "" {
				continue
			}
			MoegirlDataLines = append(MoegirlDataLines, line)
			PointRewardLines = append(PointRewardLines, reward)
		}
		lineIdInt = 0
	} else {
		_, rewardName, _ := MoegirlPointRewardBreakdown(MoegirlDataLines[lineIdInt])
		MoegirlSetItem(rewardName, data)
	}
	for lineIdInt < len(PointRewardLines) {
		PointRewardLines[lineIdInt] = MoegirlGetPointReward(MoegirlDataLines[lineIdInt])
		if PointRewardLines[lineIdInt] == "unknown_reward_name" {
			break
		} else {
			lineIdInt++
		}
	}
	if lineIdInt == len(PointRewardLines) {
		// done
		ctx.Header("Location", "/webui/event_marathon_dev/12")
	} else {
		// otherwise query with this index
		ctx.Header("Location", fmt.Sprintf("/webui/event_marathon_dev/11?line_id=%d", lineIdInt))
	}
	BuildRewardObjects()
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/11", EventMarathonDev11GET)
		router.AddHandler("/webui/event_marathon_dev", "POST", "/11", EventMarathonDev11POST)
	}
}
