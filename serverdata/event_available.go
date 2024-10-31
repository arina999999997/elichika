package serverdata

import (
	"elichika/utils"

	"xorm.io/xorm"
)

type EventAvailable struct {
	EventId int32 `xorm:"pk 'event_id'"`
	Order   int32 `xorm:"unique 'order'"`
}

func init() {
	addTable("s_event_available", EventAvailable{}, initEventAvailable)
}

// this is always manually handled
func initEventAvailable(session *xorm.Session) {
	events := []EventAvailable{}
	events = append(events, EventAvailable{
		EventId: 30001, // Secret Party!
		Order:   1,
	})
	events = append(events, EventAvailable{
		EventId: 30035, // Cutie☆Wonderland
		Order:   64,
	})
	_, err := session.Table("s_event_available").Insert(events)
	utils.CheckErr(err)
}
