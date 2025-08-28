package event_marathon_dev

import (
	"elichika/client"
	"elichika/generic"
	"elichika/router"
	"elichika/config"

	"time"
)

// contain mock data that will be returned by the relevant systems
var (
	TopStatus = client.EventMarathonTopStatus{
		StartAt:   time.Now().Unix() - 60*60*24,
		EndAt:     time.Now().Unix() + 60*60*24,
		ResultAt:  time.Now().Unix() + 60*60*23,
		ExpiredAt: time.Now().Unix() + 60*60*22,
		BgmAssetPath: client.SoundStruktur{
			V: generic.NewNullable[string]("bgm_0024"),
		},
		GachaMasterId: 123456789,
	}
	EventName           = map[string]string{}
	BoosterItemId int32 = 15001
)

func init() {
	if config.DeveloperMode == config.DeveloperModeEventMarathonDev {
		router.AddTemplates("./webui/event_marathon_dev/event_marathon_dev.html")
	}
}
