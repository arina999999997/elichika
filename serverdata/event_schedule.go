package serverdata

type EventScheduled struct {
	EventId int32 `xorm:"pk 'event_id'"`
}

func init() {
	addTable("s_event_scheduled", EventScheduled{}, nil)
}
