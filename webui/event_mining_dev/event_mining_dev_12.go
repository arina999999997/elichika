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

type VoltageRankingRewardGenerator = func() string

var voltageRankingRewardGenerators = []VoltageRankingRewardGenerator{}

func init() {
	voltageRankingRewardGenerators = append(voltageRankingRewardGenerators, func() string {
		// later voltage ranking
		eventId := TopStatus.EventId
		res := ""
		res += fmt.Sprintf("1,15,2%d01,1,1\n", eventId)
		res += fmt.Sprintf("1,21,21010,10,1\n")
		res += fmt.Sprintf("1,13,1800,150,1\n")
		res += fmt.Sprintf("1,28,16001,400,1\n")
		res += fmt.Sprintf("1,6,1400,8,1\n")
		res += fmt.Sprintf("2,15,2%d02,1,1\n", eventId)
		res += fmt.Sprintf("2,21,21010,10,1\n")
		res += fmt.Sprintf("2,13,1800,80,1\n")
		res += fmt.Sprintf("2,28,16001,200,1\n")
		res += fmt.Sprintf("2,6,1400,8,1\n")
		res += fmt.Sprintf("3,15,2%d03,1,1\n", eventId)
		res += fmt.Sprintf("3,21,21010,10,1\n")
		res += fmt.Sprintf("3,13,1800,80,1\n")
		res += fmt.Sprintf("3,28,16001,200,1\n")
		res += fmt.Sprintf("3,6,1400,8,1\n")
		res += fmt.Sprintf("4,15,2%d04,1,1\n", eventId)
		res += fmt.Sprintf("4,21,21010,10,1\n")
		res += fmt.Sprintf("4,13,1800,80,1\n")
		res += fmt.Sprintf("4,28,16001,200,1\n")
		res += fmt.Sprintf("4,6,1400,8,1\n")
		res += fmt.Sprintf("5,15,2%d05,1,1\n", eventId)
		res += fmt.Sprintf("5,21,21010,10,1\n")
		res += fmt.Sprintf("5,13,1800,80,1\n")
		res += fmt.Sprintf("5,28,16001,200,1\n")
		res += fmt.Sprintf("5,6,1400,8,1\n")
		res += fmt.Sprintf("6,15,2%d06,1,1\n", eventId)
		res += fmt.Sprintf("6,21,21010,10,1\n")
		res += fmt.Sprintf("6,13,1800,80,1\n")
		res += fmt.Sprintf("6,28,16001,150,1\n")
		res += fmt.Sprintf("6,6,1400,8,1\n")
		res += fmt.Sprintf("7,15,2%d07,1,1\n", eventId)
		res += fmt.Sprintf("7,21,21010,10,1\n")
		res += fmt.Sprintf("7,13,1800,80,1\n")
		res += fmt.Sprintf("7,28,16001,150,1\n")
		res += fmt.Sprintf("7,6,1400,8,1\n")
		res += fmt.Sprintf("8,15,2%d08,1,1\n", eventId)
		res += fmt.Sprintf("8,21,21010,10,1\n")
		res += fmt.Sprintf("8,13,1800,80,1\n")
		res += fmt.Sprintf("8,28,16001,150,1\n")
		res += fmt.Sprintf("8,6,1400,8,1\n")
		res += fmt.Sprintf("9,15,2%d09,1,1\n", eventId)
		res += fmt.Sprintf("9,21,21010,10,1\n")
		res += fmt.Sprintf("9,13,1800,80,1\n")
		res += fmt.Sprintf("9,28,16001,150,1\n")
		res += fmt.Sprintf("9,6,1400,8,1\n")
		res += fmt.Sprintf("10,15,2%d10,1,1\n", eventId)
		res += fmt.Sprintf("10,21,21010,10,1\n")
		res += fmt.Sprintf("10,13,1800,80,1\n")
		res += fmt.Sprintf("10,28,16001,150,1\n")
		res += fmt.Sprintf("10,6,1400,8,1\n")
		res += fmt.Sprintf("11,15,2%d11,1,1\n", eventId)
		res += fmt.Sprintf("11,21,21010,8,1\n")
		res += fmt.Sprintf("11,13,1800,80,1\n")
		res += fmt.Sprintf("11,28,16001,120,1\n")
		res += fmt.Sprintf("11,6,1400,8,1\n")
		res += fmt.Sprintf("51,15,2%d12,1,1\n", eventId)
		res += fmt.Sprintf("51,21,21010,8,1\n")
		res += fmt.Sprintf("51,13,1800,80,1\n")
		res += fmt.Sprintf("51,28,16001,120,1\n")
		res += fmt.Sprintf("51,6,1400,6,1\n")
		res += fmt.Sprintf("101,15,2%d13,1,1\n", eventId)
		res += fmt.Sprintf("101,21,21010,6,1\n")
		res += fmt.Sprintf("101,13,1800,80,1\n")
		res += fmt.Sprintf("101,28,16001,120,1\n")
		res += fmt.Sprintf("101,6,1400,6,1\n")
		res += fmt.Sprintf("301,15,2%d14,1,1\n", eventId)
		res += fmt.Sprintf("301,21,21010,6,1\n")
		res += fmt.Sprintf("301,13,1800,50,1\n")
		res += fmt.Sprintf("301,28,16001,120,1\n")
		res += fmt.Sprintf("301,6,1400,6,1\n")
		res += fmt.Sprintf("501,15,2%d15,1,2\n", eventId)
		res += fmt.Sprintf("501,21,21010,3,2\n")
		res += fmt.Sprintf("501,13,1800,50,2\n")
		res += fmt.Sprintf("501,28,16001,120,2\n")
		res += fmt.Sprintf("501,6,1400,4,2\n")
		res += fmt.Sprintf("1001,15,2%d16,1,2\n", eventId)
		res += fmt.Sprintf("1001,21,21010,3,2\n")
		res += fmt.Sprintf("1001,13,1800,50,2\n")
		res += fmt.Sprintf("1001,28,16001,90,2\n")
		res += fmt.Sprintf("1001,6,1400,4,2\n")
		res += fmt.Sprintf("2001,15,2%d17,1,2\n", eventId)
		res += fmt.Sprintf("2001,21,21010,3,2\n")
		res += fmt.Sprintf("2001,13,1800,50,2\n")
		res += fmt.Sprintf("2001,28,16001,90,2\n")
		res += fmt.Sprintf("2001,6,1400,4,2\n")
		res += fmt.Sprintf("3001,15,2%d18,1,2\n", eventId)
		res += fmt.Sprintf("3001,21,21010,3,2\n")
		res += fmt.Sprintf("3001,13,1800,50,2\n")
		res += fmt.Sprintf("3001,28,16001,90,2\n")
		res += fmt.Sprintf("3001,6,1400,4,2\n")
		res += fmt.Sprintf("4001,15,2%d19,1,2\n", eventId)
		res += fmt.Sprintf("4001,21,21010,3,2\n")
		res += fmt.Sprintf("4001,13,1800,50,2\n")
		res += fmt.Sprintf("4001,28,16001,90,2\n")
		res += fmt.Sprintf("4001,6,1400,4,2\n")
		res += fmt.Sprintf("5001,15,2%d20,1,3\n", eventId)
		res += fmt.Sprintf("5001,21,21010,3,3\n")
		res += fmt.Sprintf("5001,13,1800,40,3\n")
		res += fmt.Sprintf("5001,28,16001,90,3\n")
		res += fmt.Sprintf("5001,6,1400,4,3\n")
		res += fmt.Sprintf("6001,15,2%d21,1,3\n", eventId)
		res += fmt.Sprintf("6001,13,1800,40,3\n")
		res += fmt.Sprintf("6001,28,16001,90,3\n")
		res += fmt.Sprintf("6001,6,1400,2,3\n")
		res += fmt.Sprintf("7001,15,2%d22,1,3\n", eventId)
		res += fmt.Sprintf("7001,13,1800,40,3\n")
		res += fmt.Sprintf("7001,28,16001,90,3\n")
		res += fmt.Sprintf("7001,6,1400,2,3\n")
		res += fmt.Sprintf("8001,15,2%d23,1,3\n", eventId)
		res += fmt.Sprintf("8001,13,1800,30,3\n")
		res += fmt.Sprintf("8001,28,16001,90,3\n")
		res += fmt.Sprintf("8001,6,1400,1,3\n")
		res += fmt.Sprintf("9001,15,2%d24,1,3\n", eventId)
		res += fmt.Sprintf("9001,13,1800,30,3\n")
		res += fmt.Sprintf("9001,28,16001,60,3\n")
		res += fmt.Sprintf("9001,6,1400,1,3\n")
		res += fmt.Sprintf("10001,6,1400,1,3\n")
		return res
	})
	// voltageRankingRewardGenerators = append(voltageRankingRewardGenerators, func() string {
	// 	// earlier voltage ranking
	// 	eventId := TopStatus.EventId
	// 	res := ""
	// 	res += fmt.Sprintf("1,15,2%d01,1,1\n", eventId)
	// 	res += fmt.Sprintf("2,15,2%d02,1,1\n", eventId)
	// 	res += fmt.Sprintf("3,15,2%d03,1,1\n", eventId)
	// 	res += fmt.Sprintf("4,15,2%d04,1,1\n", eventId)
	// 	res += fmt.Sprintf("5,15,2%d05,1,1\n", eventId)
	// 	res += fmt.Sprintf("6,15,2%d06,1,1\n", eventId)
	// 	res += fmt.Sprintf("7,15,2%d07,1,1\n", eventId)
	// 	res += fmt.Sprintf("8,15,2%d08,1,1\n", eventId)
	// 	res += fmt.Sprintf("9,15,2%d09,1,1\n", eventId)
	// 	res += fmt.Sprintf("10,15,2%d10,1,1\n", eventId)
	// 	res += fmt.Sprintf("11,15,2%d11,1,1\n", eventId)
	// 	res += fmt.Sprintf("51,15,2%d12,1,1\n", eventId)
	// 	res += fmt.Sprintf("101,15,2%d13,1,1\n", eventId)
	// 	res += fmt.Sprintf("301,15,2%d14,1,1\n", eventId)
	// 	res += fmt.Sprintf("501,15,2%d15,1,2\n", eventId)
	// 	res += fmt.Sprintf("1001,15,2%d16,1,2\n", eventId)
	// 	res += fmt.Sprintf("2001,15,2%d17,1,2\n", eventId)
	// 	res += fmt.Sprintf("3001,15,2%d18,1,2\n", eventId)
	// 	res += fmt.Sprintf("4001,15,2%d19,1,2\n", eventId)
	// 	res += fmt.Sprintf("5001,15,2%d20,1,3\n", eventId)
	// 	res += fmt.Sprintf("6001,15,2%d21,1,3\n", eventId)
	// 	res += fmt.Sprintf("7001,15,2%d22,1,3\n", eventId)
	// 	res += fmt.Sprintf("8001,15,2%d23,1,3\n", eventId)
	// 	res += fmt.Sprintf("9001,15,2%d24,1,3\n", eventId)
	// 	res += fmt.Sprintf("10001,null,null,null,null")
	// 	return res
	// })
}

// point ranking reward
func EventMiningDev12GET(ctx *gin.Context) {
	// generator := ctx.DefaultQuery("generator", "-1\n")
	// generatorInt, err := strconv.Atoi(generator)
	// utils.CheckErr(err)
	// if generatorInt == -1 { // template is not chosen
	form := image_form.ImageForm{
		FormId:    "voltage_ranking_reward_form",
		DataLabel: "Voltage ranking reward generator type",
		DataId:    "voltage_ranking_reward_generator_type",
	}
	for i := range voltageRankingRewardGenerators {
		text := voltageRankingRewardGenerators[i]()
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

func EventMiningDev12POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	generatorType := form.Value["voltage_ranking_reward_generator_type"][0]
	generatorTypeInt, err := strconv.Atoi(generatorType)
	utils.CheckErr(err)
	voltageRankingRewardCsv = voltageRankingRewardGenerators[generatorTypeInt]()
	ctx.Header("Location", "/webui/event_mining_dev/13\n")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/12", EventMiningDev12GET)
		router.AddHandler("/webui/event_mining_dev", "POST", "/12", EventMiningDev12POST)
	}
}
