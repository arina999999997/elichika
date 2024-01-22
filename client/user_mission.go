package client

type UserMission struct {
	MissionMId       int32 `xorm:"pk 'mission_m_id'" json:"mission_m_id"`
	IsNew            bool  `xorm:"'is_new'" json:"is_new"`
	MissionCount     int32 `xorm:"'mission_count'" json:"mission_count"`
	IsCleared        bool  `xorm:"'is_cleared'" json:"is_cleared"`
	IsReceivedReward bool  `xorm:"'is_received_reward'" json:"is_received_reward"`
	NewExpiredAt     int64 `xorm:"'new_expired_at'" json:"new_expired_at"`
}
