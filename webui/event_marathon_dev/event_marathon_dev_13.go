package event_marathon_dev

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/image_form"

	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func EventMarathonDev13GET(ctx *gin.Context) {

	form := image_form.ImageForm{
		FormId:    "total_topic_reward_form",
		DataLabel: "Total topic reward order",
		DataId:    "total_topic_reward_data",
	}
	cards := []int32{}
	for _, item := range TopStatus.EventMarathonBonusPopupOrderCardMaterRows.Slice {
		if item.IsGacha {
			continue
		}
		cards = append(cards, item.CardMatserId)
	}
	var ur int32
	var sr []int32
	for _, card := range cards {
		rarity := (card / 100) % 100
		if rarity == enum.CardRarityTypeSRare {
			sr = append(sr, card)
		} else if rarity == enum.CardRarityTypeURare {
			ur = card
		}
	}
	urAssetPath, urAwakenAssetPath := image_form.GetCardAssetPaths(ur)
	sr0AssetPath, sr0AwakenAssetPath := image_form.GetCardAssetPaths(sr[0])
	sr1AssetPath, sr1AwakenAssetPath := image_form.GetCardAssetPaths(sr[1])
	form.AddImageByAssetPath(fmt.Sprintf("%d,%d,%d", ur, sr[0], sr[1]), urAssetPath)
	form.AddImageByAssetPath(fmt.Sprintf("%d,%d,%d", ur, sr[0], sr[1]), urAwakenAssetPath)
	form.AddImageByAssetPath(fmt.Sprintf("%d,%d,%d", ur, sr[0], sr[1]), sr0AssetPath)
	form.AddImageByAssetPath(fmt.Sprintf("%d,%d,%d", ur, sr[0], sr[1]), sr0AwakenAssetPath)
	form.AddImageByAssetPath(fmt.Sprintf("%d,%d,%d", ur, sr[0], sr[1]), sr1AssetPath)
	form.AddImageByAssetPath(fmt.Sprintf("%d,%d,%d", ur, sr[0], sr[1]), sr1AwakenAssetPath)
	form.AddLinebreak()
	form.AddImageByAssetPath(fmt.Sprintf("%d,%d,%d", ur, sr[1], sr[0]), urAssetPath)
	form.AddImageByAssetPath(fmt.Sprintf("%d,%d,%d", ur, sr[1], sr[0]), urAwakenAssetPath)
	form.AddImageByAssetPath(fmt.Sprintf("%d,%d,%d", ur, sr[1], sr[0]), sr1AssetPath)
	form.AddImageByAssetPath(fmt.Sprintf("%d,%d,%d", ur, sr[1], sr[0]), sr1AwakenAssetPath)
	form.AddImageByAssetPath(fmt.Sprintf("%d,%d,%d", ur, sr[1], sr[0]), sr0AssetPath)
	form.AddImageByAssetPath(fmt.Sprintf("%d,%d,%d", ur, sr[1], sr[0]), sr0AwakenAssetPath)

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_marathon_dev.html", gin.H{
		"body": msg,
	})

}

func EventMarathonDev13POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	data := form.Value["total_topic_reward_data"][0]
	tokens := strings.Split(data, ",")
	ur, err := strconv.Atoi(tokens[0])
	utils.CheckErr(err)
	sr0, err := strconv.Atoi(tokens[1])
	utils.CheckErr(err)
	sr1, err := strconv.Atoi(tokens[2])
	utils.CheckErr(err)
	TopStatus.EventTotalTopicRewardInfo.Slice = nil
	TopStatus.EventTotalTopicRewardInfo.Append(client.EventTopicReward{
		DisplayOrder: 1,
		RewardContent: client.Content{
			ContentType:   enum.ContentTypeCard,
			ContentId:     int32(ur),
			ContentAmount: 1,
		},
		// the texturestruckture are the names of the character, no need to fill it
	})
	TopStatus.EventTotalTopicRewardInfo.Append(client.EventTopicReward{
		DisplayOrder: 2,
		RewardContent: client.Content{
			ContentType:   enum.ContentTypeCard,
			ContentId:     int32(sr0),
			ContentAmount: 1,
		},
	})
	TopStatus.EventTotalTopicRewardInfo.Append(client.EventTopicReward{
		DisplayOrder: 3,
		RewardContent: client.Content{
			ContentType:   enum.ContentTypeCard,
			ContentId:     int32(sr1),
			ContentAmount: 1,
		},
	})
	ctx.Header("Location", "/webui/event_marathon_dev/14")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/13", EventMarathonDev13GET)
		router.AddHandler("/webui/event_marathon_dev", "POST", "/13", EventMarathonDev13POST)
	}
}
