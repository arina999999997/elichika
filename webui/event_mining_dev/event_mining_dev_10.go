//go:build dev

package event_mining_dev

import (
	"elichika/config"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/form/image_form"

	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type PointRankingRewardGenerator = func() string

var pointRankingRewardGenerators = []PointRankingRewardGenerator{}

func init() {
	pointRankingRewardGenerators = append(pointRankingRewardGenerators, func() string {
		eventId := TopStatus.EventId
		urCardMasterId := TopStatus.EventPointRankingTopicRewardInfo.Slice[0].RewardContent.ContentId
		srCardMasterId := TopStatus.EventPointRankingTopicRewardInfo.Slice[2].RewardContent.ContentId
		res := ""
		res += fmt.Sprintf("1,15,1%d01,1,1\n", eventId)
		res += fmt.Sprintf("1,3,%d,3,1\n", urCardMasterId)
		res += fmt.Sprintf("1,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("1,13,1800,10,1\n")
		res += fmt.Sprintf("1,12,1700,50,1\n")
		res += fmt.Sprintf("2,15,1%d02,1,1\n", eventId)
		res += fmt.Sprintf("2,3,%d,3,1\n", urCardMasterId)
		res += fmt.Sprintf("2,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("2,13,1800,10,1\n")
		res += fmt.Sprintf("2,12,1700,50,1\n")
		res += fmt.Sprintf("3,15,1%d03,1,1\n", eventId)
		res += fmt.Sprintf("3,3,%d,3,1\n", urCardMasterId)
		res += fmt.Sprintf("3,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("3,13,1800,10,1\n")
		res += fmt.Sprintf("3,12,1700,50,1\n")
		res += fmt.Sprintf("4,15,1%d04,1,1\n", eventId)
		res += fmt.Sprintf("4,3,%d,3,1\n", urCardMasterId)
		res += fmt.Sprintf("4,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("4,13,1800,10,1\n")
		res += fmt.Sprintf("4,12,1700,50,1\n")
		res += fmt.Sprintf("5,15,1%d05,1,1\n", eventId)
		res += fmt.Sprintf("5,3,%d,3,1\n", urCardMasterId)
		res += fmt.Sprintf("5,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("5,13,1800,10,1\n")
		res += fmt.Sprintf("5,12,1700,50,1\n")
		res += fmt.Sprintf("6,15,1%d06,1,1\n", eventId)
		res += fmt.Sprintf("6,3,%d,3,1\n", urCardMasterId)
		res += fmt.Sprintf("6,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("6,13,1800,10,1\n")
		res += fmt.Sprintf("6,12,1700,50,1\n")
		res += fmt.Sprintf("7,15,1%d07,1,1\n", eventId)
		res += fmt.Sprintf("7,3,%d,3,1\n", urCardMasterId)
		res += fmt.Sprintf("7,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("7,13,1800,10,1\n")
		res += fmt.Sprintf("7,12,1700,50,1\n")
		res += fmt.Sprintf("8,15,1%d08,1,1\n", eventId)
		res += fmt.Sprintf("8,3,%d,3,1\n", urCardMasterId)
		res += fmt.Sprintf("8,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("8,13,1800,10,1\n")
		res += fmt.Sprintf("8,12,1700,50,1\n")
		res += fmt.Sprintf("9,15,1%d09,1,1\n", eventId)
		res += fmt.Sprintf("9,3,%d,3,1\n", urCardMasterId)
		res += fmt.Sprintf("9,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("9,13,1800,10,1\n")
		res += fmt.Sprintf("9,12,1700,50,1\n")
		res += fmt.Sprintf("10,15,1%d10,1,1\n", eventId)
		res += fmt.Sprintf("10,3,%d,3,1\n", urCardMasterId)
		res += fmt.Sprintf("10,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("10,13,1800,10,1\n")
		res += fmt.Sprintf("10,12,1700,50,1\n")
		res += fmt.Sprintf("11,15,1%d11,1,1\n", eventId)
		res += fmt.Sprintf("11,3,%d,3,1\n", urCardMasterId)
		res += fmt.Sprintf("11,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("11,13,1800,10,1\n")
		res += fmt.Sprintf("11,12,1700,40,1\n")
		res += fmt.Sprintf("51,15,1%d12,1,1\n", eventId)
		res += fmt.Sprintf("51,3,%d,3,1\n", urCardMasterId)
		res += fmt.Sprintf("51,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("51,13,1800,10,1\n")
		res += fmt.Sprintf("51,12,1700,40,1\n")
		res += fmt.Sprintf("101,15,1%d13,1,1\n", eventId)
		res += fmt.Sprintf("101,3,%d,2,1\n", urCardMasterId)
		res += fmt.Sprintf("101,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("101,13,1800,10,1\n")
		res += fmt.Sprintf("101,12,1700,40,1\n")
		res += fmt.Sprintf("301,15,1%d14,1,1\n", eventId)
		res += fmt.Sprintf("301,3,%d,2,1\n", urCardMasterId)
		res += fmt.Sprintf("301,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("301,13,1800,10,1\n")
		res += fmt.Sprintf("301,12,1700,30,1\n")
		res += fmt.Sprintf("501,15,1%d15,1,1\n", eventId)
		res += fmt.Sprintf("501,3,%d,2,1\n", urCardMasterId)
		res += fmt.Sprintf("501,3,%d,5,1\n", srCardMasterId)
		res += fmt.Sprintf("501,13,1800,10,1\n")
		res += fmt.Sprintf("501,12,1700,30,1\n")
		res += fmt.Sprintf("1001,15,1%d16,1,2\n", eventId)
		res += fmt.Sprintf("1001,3,%d,2,2\n", urCardMasterId)
		res += fmt.Sprintf("1001,3,%d,4,2\n", srCardMasterId)
		res += fmt.Sprintf("1001,13,1800,8,2\n")
		res += fmt.Sprintf("1001,12,1700,30,2\n")
		res += fmt.Sprintf("2001,15,1%d17,1,2\n", eventId)
		res += fmt.Sprintf("2001,3,%d,2,2\n", urCardMasterId)
		res += fmt.Sprintf("2001,3,%d,4,2\n", srCardMasterId)
		res += fmt.Sprintf("2001,13,1800,8,2\n")
		res += fmt.Sprintf("2001,12,1700,20,2\n")
		res += fmt.Sprintf("3001,15,1%d18,1,2\n", eventId)
		res += fmt.Sprintf("3001,3,%d,2,2\n", urCardMasterId)
		res += fmt.Sprintf("3001,3,%d,4,2\n", srCardMasterId)
		res += fmt.Sprintf("3001,13,1800,8,2\n")
		res += fmt.Sprintf("3001,12,1700,20,2\n")
		res += fmt.Sprintf("4001,15,1%d19,1,2\n", eventId)
		res += fmt.Sprintf("4001,3,%d,2,2\n", urCardMasterId)
		res += fmt.Sprintf("4001,3,%d,4,2\n", srCardMasterId)
		res += fmt.Sprintf("4001,13,1800,8,2\n")
		res += fmt.Sprintf("4001,12,1700,20,2\n")
		res += fmt.Sprintf("5001,15,1%d20,1,2\n", eventId)
		res += fmt.Sprintf("5001,3,%d,2,2\n", urCardMasterId)
		res += fmt.Sprintf("5001,3,%d,4,2\n", srCardMasterId)
		res += fmt.Sprintf("5001,13,1800,6,2\n")
		res += fmt.Sprintf("5001,12,1700,20,2\n")
		res += fmt.Sprintf("6001,15,1%d21,1,2\n", eventId)
		res += fmt.Sprintf("6001,3,%d,2,2\n", urCardMasterId)
		res += fmt.Sprintf("6001,3,%d,4,2\n", srCardMasterId)
		res += fmt.Sprintf("6001,13,1800,6,2\n")
		res += fmt.Sprintf("6001,12,1700,20,2\n")
		res += fmt.Sprintf("7001,15,1%d22,1,2\n", eventId)
		res += fmt.Sprintf("7001,3,%d,2,2\n", urCardMasterId)
		res += fmt.Sprintf("7001,3,%d,4,2\n", srCardMasterId)
		res += fmt.Sprintf("7001,13,1800,6,2\n")
		res += fmt.Sprintf("7001,12,1700,20,2\n")
		res += fmt.Sprintf("8001,15,1%d23,1,2\n", eventId)
		res += fmt.Sprintf("8001,3,%d,2,2\n", urCardMasterId)
		res += fmt.Sprintf("8001,3,%d,4,2\n", srCardMasterId)
		res += fmt.Sprintf("8001,13,1800,6,2\n")
		res += fmt.Sprintf("8001,12,1700,20,2\n")
		res += fmt.Sprintf("9001,15,1%d24,1,2\n", eventId)
		res += fmt.Sprintf("9001,3,%d,2,2\n", urCardMasterId)
		res += fmt.Sprintf("9001,3,%d,4,2\n", srCardMasterId)
		res += fmt.Sprintf("9001,13,1800,6,2\n")
		res += fmt.Sprintf("9001,12,1700,20,2\n")
		res += fmt.Sprintf("10001,15,1%d25,1,3\n", eventId)
		res += fmt.Sprintf("10001,3,%d,1,3\n", urCardMasterId)
		res += fmt.Sprintf("10001,3,%d,3,3\n", srCardMasterId)
		res += fmt.Sprintf("10001,13,1800,4,3\n")
		res += fmt.Sprintf("10001,12,1700,10,3\n")
		res += fmt.Sprintf("20001,15,1%d26,1,3\n", eventId)
		res += fmt.Sprintf("20001,3,%d,1,3\n", urCardMasterId)
		res += fmt.Sprintf("20001,3,%d,3,3\n", srCardMasterId)
		res += fmt.Sprintf("20001,13,1800,3,3\n")
		res += fmt.Sprintf("20001,12,1700,10,3\n")
		res += fmt.Sprintf("30001,15,1%d27,1,3\n", eventId)
		res += fmt.Sprintf("30001,3,%d,1,3\n", urCardMasterId)
		res += fmt.Sprintf("30001,3,%d,3,3\n", srCardMasterId)
		res += fmt.Sprintf("30001,13,1800,2,3\n")
		res += fmt.Sprintf("30001,12,1700,10,3\n")
		res += fmt.Sprintf("40001,15,1%d28,1,3\n", eventId)
		res += fmt.Sprintf("40001,3,%d,3,3\n", srCardMasterId)
		res += fmt.Sprintf("40001,13,1800,2,3\n")
		res += fmt.Sprintf("40001,12,1700,10,3\n")
		res += fmt.Sprintf("50001,15,1%d29,1,3\n", eventId)
		res += fmt.Sprintf("50001,3,%d,2,3\n", srCardMasterId)
		res += fmt.Sprintf("50001,13,1800,1,3\n")
		res += fmt.Sprintf("50001,12,1700,5,3\n")
		res += fmt.Sprintf("60001,15,1%d30,1,3\n", eventId)
		res += fmt.Sprintf("60001,3,%d,2,3\n", srCardMasterId)
		res += fmt.Sprintf("60001,13,1800,1,3\n")
		res += fmt.Sprintf("60001,12,1700,5,3\n")
		res += fmt.Sprintf("70001,15,1%d31,1,3\n", eventId)
		res += fmt.Sprintf("70001,3,%d,1,3\n", srCardMasterId)
		res += fmt.Sprintf("70001,13,1800,1,3\n")
		res += fmt.Sprintf("70001,12,1700,5,3\n")
		res += fmt.Sprintf("80001,15,1%d32,1,3\n", eventId)
		res += fmt.Sprintf("80001,3,%d,1,3\n", srCardMasterId)
		res += fmt.Sprintf("80001,13,1800,1,3\n")
		res += fmt.Sprintf("80001,12,1700,5,3\n")
		res += fmt.Sprintf("90001,15,1%d33,1,3\n", eventId)
		res += fmt.Sprintf("90001,3,%d,1,3\n", srCardMasterId)
		res += fmt.Sprintf("90001,13,1800,1,3\n")
		res += fmt.Sprintf("90001,12,1700,5,3\n")
		res += fmt.Sprintf("100001,15,1%d98,1,3\n", eventId)
		res += fmt.Sprintf("100001,13,1800,1,3\n")
		res += fmt.Sprintf("100001,12,1700,5,3\n")
		res += fmt.Sprintf("150001,15,1%d98,1,3\n", eventId)
		res += fmt.Sprintf("150001,12,1700,5,3\n")
		res += fmt.Sprintf("200001,15,1%d98,1,3\n", eventId)
		return res
	})
}

// point ranking reward
func EventMiningDev10GET(ctx *gin.Context) {
	// generator := ctx.DefaultQuery("generator", "-1")
	// generatorInt, err := strconv.Atoi(generator)
	// utils.CheckErr(err)
	// if generatorInt == -1 { // template is not chosen
	form := image_form.ImageForm{
		FormId:    "point_ranking_reward_form",
		DataLabel: "Point ranking reward generator type",
		DataId:    "point_ranking_reward_generator_type",
	}
	for i := range pointRankingRewardGenerators {
		text := pointRankingRewardGenerators[i]()
		// generate a few reward line
		lines := strings.Split(text, "\n")
		for j := 0; j < 10; j++ {
			content := parseContent(lines[j])
			form.AddImageByContent(fmt.Sprint(i), content)
			if j%5 == 4 {
				form.AddLinebreak()
			}
		}
		form.AddLinebreak()
	}
	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_mining_dev.html", gin.H{
		"body": msg,
	})
	// }

	// ctx.Header("Content-Type", "text/html")
	// msg := form.GetHTML()
	// ctx.HTML(http.StatusOK, "event_mining_dev.html", gin.H{
	// 	"body": msg,
	// })

}

func EventMiningDev10POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	generatorType := form.Value["point_ranking_reward_generator_type"][0]
	generatorTypeInt, err := strconv.Atoi(generatorType)
	utils.CheckErr(err)
	pointRankingRewardCsv = pointRankingRewardGenerators[generatorTypeInt]()
	ctx.Header("Location", "/webui/event_mining_dev/11")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/10", EventMiningDev10GET)
		router.AddHandler("/webui/event_mining_dev", "POST", "/10", EventMiningDev10POST)
	}
}
