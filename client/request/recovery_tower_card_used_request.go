package request

import (
	"elichika/generic"
)

type RecoveryTowerCardUsedRequest struct {
	TowerId       int32                `json:"tower_id"`
	CardMasterIds generic.Array[int32] `json:"card_master_ids"`
}
