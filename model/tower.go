package model

type UserTower struct {
	UserID                      int `xorm:"pk 'user_id'" json:"-"`
	TowerID                     int `xorm:"pk 'tower_id'" json:"tower_id"`
	ClearedFloor                int `xorm:"'cleared_floor'" json:"cleared_floor"`
	ReadFloor                   int `xorm:"'read_floor'" json:"read_floor"`
	Voltage                     int `xorm:"'voltage'" json:"voltage"`
	RecoveryPointFullAt         int `xorm:"'recovery_point_full_at'" json:"recovery_point_full_at"`
	RecoveryPointLastConsumedAt int `xorm:"'recovery_point_last_consumed_at'" json:"recovery_point_last_consumed_at"`
}

func (ut *UserTower) ID() int64 {
	return int64(ut.TowerID)
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_tower"] = UserTower{}
}
