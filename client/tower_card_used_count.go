package client

type TowerCardUsedCount struct {
	CardMasterId   int32 `xorm:"pk" json:"card_master_id"`
	UsedCount      int32 `json:"used_count"`
	RecoveredCount int32 `json:"recovered_count"`
	LastUsedAt     int64 `json:"last_used_at"`
}
