package client

type LiveTowerStatus struct {
	TowerId int32 `xorm:"'tower_id'" json:"tower_id"`
	FloorNo int32 `xorm:"'floor_no'" json:"floor_no"`
}
