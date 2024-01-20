package client

import (
	"elichika/generic"
)

type LiveResultTower struct {
	TowerId             int32                             `json:"tower_id"`
	FloorNo             int32                             `json:"floor_no"`
	TotalVoltage        int32                             `json:"total_voltage"`
	GettedVoltage       int32                             `json:"getted_voltage"`
	TowerCardUsedCounts generic.Array[TowerCardUsedCount] `json:"tower_card_used_counts"`
}
