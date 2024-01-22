package client

import (
	"elichika/generic"
)

type UserWeeklyMission struct {
	MissionMId        int32                   `xorm:"pk 'mission_m_id'" json:"mission_m_id"`
	IsNew             bool                    `xorm:"'is_new'" json:"is_new"`
	MissionStartCount int32                   `xorm:"'mission_start_count'" json:"mission_start_count"`
	MissionCount      int32                   `xorm:"'mission_count'" json:"mission_count"`
	IsCleared         bool                    `xorm:"'is_cleared'" json:"is_cleared"`
	IsReceivedReward  bool                    `xorm:"'is_received_reward'" json:"is_received_reward"`
	ClearedExpiredAt  generic.Nullable[int64] `xorm:"json 'cleared_expired_at'" json:"cleared_expired_at"`
	NewExpiredAt      int64                   `xorm:"'new_expired_at'" json:"new_expired_at"`
}
