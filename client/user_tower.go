package client

type UserTower struct {
	TowerId                     int32 `xorm:"pk 'tower_id'" json:"tower_id"`
	ClearedFloor                int32 `xorm:"'cleared_floor'" json:"cleared_floor"`
	ReadFloor                   int32 `xorm:"'read_floor'" json:"read_floor"`
	Voltage                     int32 `xorm:"'voltage'" json:"voltage"`
	RecoveryPointFullAt         int64 `xorm:"'recovery_point_full_at'" json:"recovery_point_full_at"`
	RecoveryPointLastConsumedAt int64 `xorm:"'recovery_point_last_consumed_at'" json:"recovery_point_last_consumed_at"`
}
func (ut *UserTower) Id() int64 {
	return int64(ut.TowerId)
}