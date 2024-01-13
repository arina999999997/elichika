package client

import (
	"elichika/generic"
)

type IllustLoginBonus struct {
	LoginBonusId      int32                           `json:"login_bonus_id"`
	LoginBonusRewards generic.List[LoginBonusRewards] `json:"login_bonus_rewards"`
	BackgroundId      int32                           `json:"background_id"`
	StartAt           int64                           `json:"start_at"`
	EndAt             int64                           `json:"end_at"`
}
