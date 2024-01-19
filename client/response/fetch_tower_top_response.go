package response

import (
	"elichika/client"
	"elichika/generic"
)

type FetchTowerTopResponse struct {
	TowerCardUsedCountRows generic.List[client.TowerCardUsedCount] `json:"tower_card_used_count_rows"`
	IsShowUnlockEffect     bool                                    `json:"is_show_unlock_effect"`
	UserModelDiff          *client.UserModel                       `json:"user_model_diff"`
	Order                  generic.Nullable[int32]                 `json:"order"`
	EachBonusLiveVoltage   generic.List[int32]                     `json:"each_bonus_live_voltage"`
}
