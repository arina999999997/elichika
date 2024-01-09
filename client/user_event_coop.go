package client

type UserEventCoop struct {
	EventMasterId     int32 `xorm:"pk 'event_master_id'" json:"event_master_id"`
	CurrentRoomId     int32 `xorm:"pk 'current_room_id'" json:"current_room_id"`
	EventPoint        int32 `xorm:"'event_point'" json:"event_point"`
	RecentAwardId     int32 `xorm:"'recent_award_id'" json:"recent_award_id"`
	EventVoltagePoint int32 `xorm:"'event_voltage_point'" json:"event_voltage_point"`
	CoopPoint         int32 `xorm:"'coop_point'" json:"coop_point"`
	CoopPointResetAt  int64 `xorm:"'coop_point_reset_at'" json:"coop_point_reset_at"`
	CoopPointBroken   int32 `xorm:"'coop_point_broken'" json:"coop_point_broken"`
	PlayableAt        int64 `xorm:"'playable_at'" json:"playable_at"`
	PenaltyCount      int32 `xorm:"'penalty_count'" json:"penalty_count"`
}

func (uec *UserEventCoop) Id() int64 {
	return int64(uec.EventMasterId)
}
