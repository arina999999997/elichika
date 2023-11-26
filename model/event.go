package model

// TODO: not actually implemented, just for archival purpose
type UserEventMarathon struct { // Story ranking event
	UserID             int `xorm:"pk 'user_id'" json:"-"`
	EventMasterID      int `xorm:"pk 'event_master_id'" json:"event_master_id"`
	EventPoint         int `xorm:"'event_point'" json:"event_point"`
	OpennedStoryNumber int `xorm:"'opened_story_number'" json:"opened_story_number"`
	ReadStoryNumber    int `xorm:"'read_story_number'" json:"read_story_number"`
}

func (uem *UserEventMarathon) ID() int64 {
	return int64(uem.EventMasterID)
}

type UserEventMining struct { // Voltage ranking event
	UserID             int `xorm:"pk 'user_id'" json:"-"`
	EventMasterID      int `xorm:"pk 'event_master_id'" json:"event_master_id"`
	EventPoint         int `xorm:"'event_point'" json:"event_point"`
	EventVoltagePoint  int `xorm:"'event_voltage_point'" json:"event_voltage_point"`
	OpennedStoryNumber int `xorm:"'opened_story_number'" json:"opened_story_number"`
	ReadStoryNumber    int `xorm:"'read_story_number'" json:"read_story_number"`
}

func (uem *UserEventMining) ID() int64 {
	return int64(uem.EventMasterID)
}

type UserEventCoop struct { // SBL
	UserID            int `xorm:"pk 'user_id'" json:"-"`
	EventMasterID     int `xorm:"pk 'event_master_id'" json:"event_master_id"`
	CurrentRoomID     int `xorm:"pk 'current_room_id'" json:"current_room_id"`
	EventPoint        int `xorm:"'event_point'" json:"event_point"`
	RecentAwardID     int `xorm:"'recent_award_id'" json:"recent_award_id"`
	EventVoltagePoint int `xorm:"'event_voltage_point'" json:"event_voltage_point"`
	CoopPoint         int `xorm:"'coop_point'" json:"coop_point"`
	CoopPointResetAt  int `xorm:"'coop_point_reset_at'" json:"coop_point_reset_at"`
	CoopPointBroken   int `xorm:"'coop_point_broken'" json:"coop_point_broken"`
	PlayableAt        int `xorm:"'playable_at'" json:"playable_at"`
	PenaltyCount      int `xorm:"'penalty_count'" json:"penalty_count"`
}

func (uec *UserEventCoop) ID() int64 {
	return int64(uec.EventMasterID)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_event_marathon"] = UserEventMarathon{}
	TableNameToInterface["u_event_mining"] = UserEventMining{}
	TableNameToInterface["u_event_coop"] = UserEventCoop{}
}
