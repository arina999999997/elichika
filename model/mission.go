package model

import (
	"elichika/generic"
)

// TODO: saved in database but not handled

type UserMission struct {
	MissionMId       int  `xorm:"pk 'mission_m_id'" json:"mission_m_id"`
	IsNew            bool `xorm:"'is_new'" json:"is_new"`
	MissionCount     int  `xorm:"'mission_count'" json:"mission_count"`
	IsCleared        bool `xorm:"'is_cleared'" json:"is_cleared"`
	IsReceivedReward bool `xorm:"'is_received_reward'" json:"is_received_reward"`
	NewExpiredAt     int  `xorm:"'new_expired_at'" json:"new_expired_at"`
}

func (um *UserMission) Id() int64 {
	return int64(um.MissionMId)
}

type UserDailyMission struct {
	MissionMId        int  `xorm:"pk 'mission_m_id'" json:"mission_m_id"`
	IsNew             bool `xorm:"'is_new'" json:"is_new"`
	MissionStartCount int  `xorm:"'mission_start_count'" json:"mission_start_count"`
	MissionCount      int  `xorm:"'mission_count'" json:"mission_count"`
	IsCleared         bool `xorm:"'is_cleared'" json:"is_cleared"`
	IsReceivedReward  bool `xorm:"'is_received_reward'" json:"is_received_reward"`
	ClearedExpiredAt  int  `xorm:"'cleared_expired_at'" json:"cleared_expired_at"`
}

func (udm *UserDailyMission) Id() int64 {
	return int64(udm.MissionMId)
}

type UserWeeklyMission struct {
	MissionMId        int  `xorm:"pk 'mission_m_id'" json:"mission_m_id"`
	IsNew             bool `xorm:"'is_new'" json:"is_new"`
	MissionStartCount int  `xorm:"'mission_start_count'" json:"mission_start_count"`
	MissionCount      int  `xorm:"'mission_count'" json:"mission_count"`
	IsCleared         bool `xorm:"'is_cleared'" json:"is_cleared"`
	IsReceivedReward  bool `xorm:"'is_received_reward'" json:"is_received_reward"`
	ClearedExpiredAt  int  `xorm:"'cleared_expired_at'" json:"cleared_expired_at"`
	NewExpiredAt      int  `xorm:"'new_expired_at'" json:"new_expired_at"`
}

func (uwm *UserWeeklyMission) Id() int64 {
	return int64(uwm.MissionMId)
}

func init() {

	TableNameToInterface["u_mission"] = generic.UserIdWrapper[UserMission]{}
	TableNameToInterface["u_daily_mission"] = generic.UserIdWrapper[UserDailyMission]{}
	TableNameToInterface["u_weekly_mission"] = generic.UserIdWrapper[UserWeeklyMission]{}
}
