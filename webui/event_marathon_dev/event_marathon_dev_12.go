//go:build dev

package event_marathon_dev

import (
	"elichika/config"
	"elichika/router"
	"elichika/utils"
	// "elichika/client"
	"elichika/enum"

	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	RankingRewardLines = []string{}
)

// This is where we build the reward for the event
// TODO: for now only support early moegirl, probably should select the method first if we have method
func EventMarathonDev12GET(ctx *gin.Context) {
	msg := `<form id="ranking_reward_form" method="POST" enctype="multipart/form-data">`
	msg += "\n"
	msg += `<div><label>Upload moegirl old format ranking file: <input type="file" name="file"/> <input type="text" name="ranking_reward_data"/>`
	msg += "\n"
	ctx.Header("Content-Type", "text/html")
	ctx.HTML(http.StatusOK, "event_marathon_dev.html", gin.H{
		"body": msg,
	})
}

func EventMarathonDev12POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	// data := form.Value["ranking_reward_data"][0]
	fmt.Println(form)

	fileHeader := form.File["file"][0]
	file, err := fileHeader.Open()
	utils.CheckErr(err)
	bytes := make([]byte, fileHeader.Size)
	_, err = file.Read(bytes)
	utils.CheckErr(err)
	lines := strings.Split(string(bytes), "\r\n")
	rewards := []string{"rank col"}
	// all events with the exception of the first one use title id as:
	// 10000000 + eventId * 100 + i
	// i = 1 -> 33 for the numbered one, i = 98 for the participation award
	// the first event use i = 34 for partitipation award
	rewardTierId := 0
	// same format used for ranking reward.csv
	// highest_rank,content_type,content_id,content_amount,ranking_result_prize_type
	RankingRewardLines = []string{}
	for i, line := range lines {
		if line == "" {
			continue
		}
		if i == 0 {
			continue
		}
		if i == 2 {
			// parse the rewards
			rewardTokens := strings.Split(line, "\t")
			for j := range rewardTokens {
				rewardTokens[j] = strings.TrimSpace(rewardTokens[j])
				if j == 0 {
					rewards = append(rewards, "title")
				} else {
					rewards = append(rewards, MoegirlMustGetItem(rewardTokens[j]))
				}
			}
			continue
		}
		// reward tier
		rewardTierId += 1
		titleId := 10000000 + int(TopStatus.EventId)*100
		if rewardTierId <= 33 {
			titleId += rewardTierId
		} else {
			titleId += 98
		}
		//
		rewardTokens := strings.Split(line, "\t")
		for j := range rewardTokens {
			rewardTokens[j] = strings.TrimSpace(rewardTokens[j])
		}
		highestRank := MoegirlParseHighestRank(rewardTokens[0])
		// title is known
		prizeType := enum.EventRankingResultPrizeTypeGold
		if highestRank > 1000 {
			prizeType = enum.EventRankingResultPrizeTypeSilver
		}
		if highestRank > 10000 {
			prizeType = enum.EventRankingResultPrizeTypeBronze
		}
		// title
		RankingRewardLines = append(RankingRewardLines, fmt.Sprintf("%d,%d,%d,%d,%d", highestRank, enum.ContentTypeEmblem, titleId, 1, prizeType))
		// first card
		for j := 2; j < len(rewardTokens); j++ {
			if rewardTokens[j] == "" {
				continue
			}
			RankingRewardLines = append(RankingRewardLines, fmt.Sprintf("%d,%s,%d,%d", highestRank, rewards[j], MoegirlParseCount(rewardTokens[j]), prizeType))
		}
		if highestRank == 200001 {
			break
		}
	}
	BuildRewardObjects()
	ctx.Header("Location", "/webui/event_marathon_dev/13")
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddHandler("/webui/event_marathon_dev", "GET", "/12", EventMarathonDev12GET)
		router.AddHandler("/webui/event_marathon_dev", "POST", "/12", EventMarathonDev12POST)
	}
}
