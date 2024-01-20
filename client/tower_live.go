package client

type TowerLive struct {
	TowerId       int32 `json:"tower_id"`
	FloorNo       int32 `json:"floor_no"`
	TargetVoltage int32 `json:"target_voltage"`
	StartVoltage  int32 `json:"start_voltage"`
}
