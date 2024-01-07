package model

type UserTower struct {
	UserId                      int `xorm:"pk 'user_id'" json:"-"`
	TowerId                     int `xorm:"pk 'tower_id'" json:"tower_id"`
	ClearedFloor                int `xorm:"'cleared_floor'" json:"cleared_floor"`
	ReadFloor                   int `xorm:"'read_floor'" json:"read_floor"`
	Voltage                     int `xorm:"'voltage'" json:"voltage"`
	RecoveryPointFullAt         int `xorm:"'recovery_point_full_at'" json:"recovery_point_full_at"`
	RecoveryPointLastConsumedAt int `xorm:"'recovery_point_last_consumed_at'" json:"recovery_point_last_consumed_at"`
}

func (ut *UserTower) Id() int64 {
	return int64(ut.TowerId)
}

type UserTowerCardUsedCount struct {
	UserId         int   `xorm:"pk 'user_id'" json:"-"`
	TowerId        int   `xorm:"pk 'tower_id'" json:"-"`
	CardMasterId   int   `xorm:"pk 'card_master_id'" json:"card_master_id"`
	UsedCount      int   `xorm:"'used_count'" json:"used_count"`
	RecoveredCount int   `xorm:"'recovered_count'" json:"recovered_count"`
	LastUsedAt     int64 `xorm:"'last_used_at'" json:"last_used_at"`
}

type UserTowerVoltageRankingScore struct {
	UserId  int `xorm:"pk 'user_id'"`
	TowerId int `xorm:"pk 'tower_id'"`
	FloorNo int `xorm:"pk 'floor_no'"`
	Voltage int `xorm:"'voltage'"`
}

type LiveTowerStatus struct {
	TowerId int `xorm:"'tower_id'" json:"tower_id"`
	FloorNo int `xorm:"'floor_no'" json:"floor_no"`
}

type TowerLive struct {
	TowerId       *int `xorm:"tower_id" json:"tower_id"`
	FloorNo       *int `xorm:"tower_floor_no" json:"floor_no"`
	TargetVoltage *int `xorm:"tower_target_voltage" json:"target_voltage"`
	StartVoltage  *int `xorm:"tower_start_voltage" json:"start_voltage"`
}

func init() {
	if TableNameToInterface == nil {
		TableNameToInterface = make(map[string]interface{})
	}
	TableNameToInterface["u_tower"] = UserTower{}
	TableNameToInterface["u_tower_card_used"] = UserTowerCardUsedCount{}
	TableNameToInterface["u_tower_voltage_ranking_score"] = UserTowerVoltageRankingScore{}
}
