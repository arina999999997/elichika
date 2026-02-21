//go:build dev

package event_mining_dev

import (
	"elichika/client"
	"elichika/config"
	"elichika/enum"
	"elichika/gamedata"
	"elichika/router"
	"elichika/utils"
	"elichika/webui/form/select_form"

	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// song is done by name selection

func EventMiningDev06GET(ctx *gin.Context) {
	liveNumber := ctx.DefaultQuery("live", "1")
	form := select_form.SelectForm{
		FormId:    "event_top_live_id",
		DataLabel: "Event top still lives (left to right)",
		DataId:    "event_top_live_id_" + fmt.Sprint(liveNumber),
	}
	gamedata := gamedata.GamedataByLocale["en"]
	liveIds := []int32{}
	for liveId := range gamedata.Live {
		if (liveId < 10000) || (liveId >= 20000) {
			continue
		}
		goodLiveDiff := false
		for _, live := range gamedata.Live[liveId].LiveDifficulties {
			goodLiveDiff = goodLiveDiff || (live.UnlockPattern == enum.LiveUnlockPatternOpen);
			goodLiveDiff = goodLiveDiff || (live.UnlockPattern == enum.LiveUnlockPatternMemberStory);
			goodLiveDiff = goodLiveDiff || (live.UnlockPattern == enum.LiveUnlockPatternExtra);
		}
		if goodLiveDiff {
			liveIds = append(liveIds, liveId)
		}
	}
	sort.Slice(liveIds, func(i, j int) bool {
		return liveIds[i] < liveIds[j]
	})
	for _, liveId := range liveIds {
		name := gamedata.Live[liveId].Name
		name = strings.ReplaceAll(name, "&apos;", "'")
		name = strings.ReplaceAll(name, "&quot;", "\"")
		name = strings.ReplaceAll(name, "&amp;", "&")
		form.AddOption(fmt.Sprint(liveId), name)
	}

	ctx.Header("Content-Type", "text/html")
	msg := form.GetHTML()
	ctx.HTML(http.StatusOK, "event_mining_dev.html", gin.H{
		"body": msg,
	})
}

func EventMiningDev06POST(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	utils.CheckErr(err)
	fmt.Println(form.Value)
	liveNumber := ctx.DefaultQuery("live", "1")
	liveNumberInt, err := strconv.Atoi(liveNumber)
	utils.CheckErr(err)
	liveId := form.Value["event_top_live_id_"+liveNumber][0]
	liveIdInt, err := strconv.Atoi(liveId)
	utils.CheckErr(err)
	if liveNumberInt == 1 {
		TopStatus.EventMiningCompetitionMasterRows.Slice = nil
	}
	TopStatus.EventMiningCompetitionMasterRows.Append(
		client.EventMiningCompetitionMasterRow{
			LiveId: int32(liveIdInt),
		})
	if liveNumberInt == 3 { // done
		ctx.Header("Location", "/webui/event_mining_dev/07")
	} else {
		ctx.Header("Location", fmt.Sprintf("/webui/event_mining_dev/06?live=%d", liveNumberInt+1))
	}
	ctx.String(http.StatusSeeOther, "")
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddHandler("/webui/event_mining_dev", "GET", "/06", EventMiningDev06GET)
		router.AddHandler("/webui/event_mining_dev", "POST", "/06", EventMiningDev06POST)
	}
}
