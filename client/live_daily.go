package client

import (
	"elichika/generic"
)

type LiveDaily struct {
	LiveDailyMasterId      int32                   `json:"live_daily_master_id"`
	LiveMasterId           int32                   `json:"live_master_id"`
	EndAt                  int64                   `json:"end_at"`
	RemainingPlayCount     int32                   `json:"remaining_play_count"`
	RemainingRecoveryCount generic.Nullable[int32] `json:"remaining_recovery_count"`
}
