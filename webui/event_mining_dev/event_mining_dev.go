//go:build dev

package event_mining_dev

import (
	"elichika/client"
	"elichika/config"
	"elichika/generic"
	"elichika/router"

	"strconv"
	"strings"
	"time"
)

// contain mock data that will be returned by the relevant systems
var (
	TopStatus = client.EventMiningTopStatus{
		StartAt:   time.Now().Unix() - 60*60*24,
		EndAt:     time.Now().Unix() + 60*60*24,
		ResultAt:  time.Now().Unix() + 60*60*23,
		ExpiredAt: time.Now().Unix() + 60*60*22,
		BgmAssetPath: client.SoundStruktur{
			V: generic.NewNullable[string]("bgm_0023"),
		},
		GachaMasterId: 0,
	}
	pointRankingRewardCsv   string
	voltageRankingRewardCsv string
	tradeProducts           = []EventMiningTradeProduct{}
)

type EventMiningTradeProduct struct {
	ContentType   int32
	ContentId     int32
	ContentAmount int32
	Price         int32
	Stock         int32
}

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMiningDev {
		router.AddTemplates("./webui/event_mining_dev/event_mining_dev.html")
	}
}

func parseContent(rankingRewardLine string) client.Content {
	tokens := strings.Split(rankingRewardLine, ",")
	contentType, _ := strconv.Atoi(tokens[1])
	contentId, _ := strconv.Atoi(tokens[2])
	contentAmount, _ := strconv.Atoi(tokens[3])
	return client.Content{
		ContentType:   int32(contentType),
		ContentId:     int32(contentId),
		ContentAmount: int32(contentAmount),
	}
}
