package serverstate

type EventActive struct {
	EventId   int32 `xorm:"pk 'event_id'"`
	EventType int32 `xorm:"'event_type'" enum:"EventType1"`
	StartAt   int64 `xorm:"'start_at'"`
	ExpiredAt int64 `xorm:"'expired_at'"`
	ResultAt  int64 `xorm:"'result_at'"`
	EndAt     int64 `xorm:"'end_at'"`
}

func init() {
	addTable("s_event_active", EventActive{}, nil)
}
