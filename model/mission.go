package model

// TODO: saved in database but not handled

type UserMission struct {
	UserID           int  `xorm:"pk 'user_id'" json:"-"`
	MissionMID       int  `xorm:"pk 'mission_m_id'" json:"mission_m_id"`
	IsNew            bool `xorm:"'is_new'" json:"is_new"`
	MissionCount     int  `xorm:"'mission_count'" json:"mission_count"`
	IsCleared        bool `xorm:"'is_cleared'" json:"is_cleared"`
	IsReceivedReward bool `xorm:"'is_received_reward'" json:"is_received_reward"`
	NewExpiredAt     int  `xorm:"'new_expired_at'" json:"new_expired_at"`
}

func (um *UserMission) ID() int64 {
	return int64(um.MissionMID)
}

type UserDailyMission struct {
	UserID            int  `xorm:"pk 'user_id'" json:"-"`
	MissionMID        int  `xorm:"pk 'mission_m_id'" json:"mission_m_id"`
	IsNew             bool `xorm:"'is_new'" json:"is_new"`
	MissionStartCount int  `xorm:"'mission_start_count'" json:"mission_start_count"`
	MissionCount      int  `xorm:"'mission_count'" json:"mission_count"`
	IsCleared         bool `xorm:"'is_cleared'" json:"is_cleared"`
	IsReceivedReward  bool `xorm:"'is_received_reward'" json:"is_received_reward"`
	ClearedExpiredAt  int  `xorm:"'cleared_expired_at'" json:"cleared_expired_at"`
}

func (udm *UserDailyMission) ID() int64 {
	return int64(udm.MissionMID)
}

type UserWeeklyMission struct {
	UserID            int  `xorm:"pk 'user_id'" json:"-"`
	MissionMID        int  `xorm:"pk 'mission_m_id'" json:"mission_m_id"`
	IsNew             bool `xorm:"'is_new'" json:"is_new"`
	MissionStartCount int  `xorm:"'mission_start_count'" json:"mission_start_count"`
	MissionCount      int  `xorm:"'mission_count'" json:"mission_count"`
	IsCleared         bool `xorm:"'is_cleared'" json:"is_cleared"`
	IsReceivedReward  bool `xorm:"'is_received_reward'" json:"is_received_reward"`
	ClearedExpiredAt  int  `xorm:"'cleared_expired_at'" json:"cleared_expired_at"`
	NewExpiredAt      int  `xorm:"'new_expired_at'" json:"new_expired_at"`
}

func (uwm *UserWeeklyMission) ID() int64 {
	return int64(uwm.MissionMID)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_mission"] = UserMission{}
	TableNameToInterface["u_daily_mission"] = UserDailyMission{}
	TableNameToInterface["u_weekly_mission"] = UserWeeklyMission{}
}
