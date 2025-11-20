//go:build dev

package event_mining_dev

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/generic"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/form/image_form"

	"net/http"

	"github.com/gin-gonic/gin"
)

func EventMiningDev11GET(ctx *gin.Context) {
	form := image_form.ImageForm{
		FormId:    "voltage_ranking_topic_reward_form",
		DataLabel: "Voltage ranking topic reward asset path",
		DataId:    "voltage_ranking_topic_reward_data",
	}

	assetPaths := image_form.GetAssetPathsFromSamePackage(TopStatus.TitleImagePath.V.Value)
	for _, assetPath := range assetPaths {
		form.AddImageByAssetPathFilterBySize(assetPath, assetPath, 426, 364)
	}
	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_mining_dev.html", gin.H{
		"body": msg,
	})

}

func EventMiningDev11POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	assetPath := form.Value["voltage_ranking_topic_reward_data"][0]
	TopStatus.EventVoltageRankingTopicRewardInfo.Append(client.EventMiningTopicReward{
		DisplayOrder: 1,
		RewardContent: client.Content{
			ContentType:   12,
			ContentId:     1700,
			ContentAmount: 1,
		},
		RankingCategory: enum.EventRankingCategoryVoltage,
		SinglePictureAssetPath: client.TextureStruktur{
			V: generic.NewNullable(assetPath),
		},
	})
	ctx.Header("Location", "/webui/event_mining_dev/12")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/11", EventMiningDev11GET)
		router.AddHandler("/webui/event_mining_dev", "POST", "/11", EventMiningDev11POST)
	}
}
