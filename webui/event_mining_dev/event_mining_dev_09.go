//go:build dev

package event_mining_dev

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/form/image_form"

	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func EventMiningDev09GET(ctx *gin.Context) {
	form := image_form.ImageForm{
		FormId:    "point_ranking_topic_reward_form",
		DataLabel: "Point ranking topic reward order",
		DataId:    "point_ranking_topic_reward_data",
	}
	cards := []int32{}
	for _, item := range TopStatus.EventMiningBonusPopupOrderCardMaterRows.Slice {
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
	ctx.HTML(http.StatusOK, "event_mining_dev.html", gin.H{
		"body": msg,
	})

}

func EventMiningDev09POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	data := form.Value["point_ranking_topic_reward_data"][0]
	tokens := strings.Split(data, ",")
	ur, err := strconv.Atoi(tokens[0])
	utils.CheckErr(err)
	sr0, err := strconv.Atoi(tokens[1])
	utils.CheckErr(err)
	sr1, err := strconv.Atoi(tokens[2])
	utils.CheckErr(err)
	TopStatus.EventPointRankingTopicRewardInfo.Slice = nil
	TopStatus.EventPointRankingTopicRewardInfo.Append(client.EventMiningTopicReward{
		DisplayOrder: 1,
		RewardContent: client.Content{
			ContentType:   enum.ContentTypeCard,
			ContentId:     int32(ur),
			ContentAmount: 1,
		},
		RankingCategory: enum.EventRankingCategoryPoint,
		// the texturestruckture are the names of the character, no need to fill it
	})
	TopStatus.EventPointRankingTopicRewardInfo.Append(client.EventMiningTopicReward{
		DisplayOrder: 2,
		RewardContent: client.Content{
			ContentType:   enum.ContentTypeCard,
			ContentId:     int32(sr0),
			ContentAmount: 1,
		},
		RankingCategory: enum.EventRankingCategoryPoint,
	})
	TopStatus.EventPointRankingTopicRewardInfo.Append(client.EventMiningTopicReward{
		DisplayOrder: 3,
		RewardContent: client.Content{
			ContentType:   enum.ContentTypeCard,
			ContentId:     int32(sr1),
			ContentAmount: 1,
		},
		RankingCategory: enum.EventRankingCategoryPoint,
	})
	ctx.Header("Location", "/webui/event_mining_dev/10")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/09", EventMiningDev09GET)
		router.AddHandler("/webui/event_mining_dev", "POST", "/09", EventMiningDev09POST)
	}
}
