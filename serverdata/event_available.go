package serverdata

import (
	"elichika/config"
	"elichika/utils"

	"os"
	"strconv"

	"xorm.io/xorm"
)

type EventAvailable struct {
	EventId int32 `xorm:"pk 'event_id'"`
	Order   int32 `xorm:"unique 'order'"`
}

func init() {
	addTable("s_event_available", EventAvailable{}, initEventAvailable)
}

func initEventAvailable(session *xorm.Session) {
	events := []EventAvailable{}

	entries, err := os.ReadDir(config.AssetPath + "event/marathon")
	utils.CheckErr(err)
	for _, entry := range entries {
		eventId, err := strconv.Atoi(entry.Name())
		utils.CheckErr(err)
		events = append(events, EventAvailable{
			EventId: int32(eventId),
			Order:   int32(eventId),
		})
	}
	_, err = session.Table("s_event_available").Insert(events)
	utils.CheckErr(err)
}
